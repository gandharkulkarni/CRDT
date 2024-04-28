package main

import (
	"bufio"
	"crdt/src/comms_handler"
	"crdt/src/constants"
	"crdt/src/helper"
	lww "crdt/src/last_write_wins"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

type collabNode struct {
	machine    string
	listenPort int32
	talkPort   int32
}

var dataChannel chan string
var quitChannel chan os.Signal
var port *int
var lwwReg *lww.LWWRegister
var id int64 = 0

func main() {
	port = flag.Int("port", constants.COLLAB_PORT, "Listner port")
	central := flag.String("central", constants.CENTRAL, "Central machine:port")
	collaborator := flag.String("collab", "", "Other collaborator node name")

	flag.Parse()

	if *port == 0 {
		panic("Insufficient number of arguments. Usage: main.go -port=<port>")
	}
	if *central == "" {
		panic("Insufficient number of arguments. Usage: main.go -central=<machine:port>")
	}
	fmt.Println("Listener port :", *port)
	// fmt.Println("Sender port :", *port+1)
	machine, err := os.Hostname() //machine.domain
	helper.CheckErr(err)
	machine = strings.Split(machine, ".")[0]

	// Create a channel to communicate between goroutines
	dataChannel = make(chan string)

	// Create a channel to listen for a signal to quit
	quitChannel = make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	centralHostDetails := *central
	queryPort, err := strconv.Atoi(strings.Split(centralHostDetails, ":")[1])
	helper.CheckErr(err)

	queryHostDetails := strings.Split(centralHostDetails, ":")[0] + ":" + strconv.Itoa(queryPort+1)

	registerWithCentralNode(centralHostDetails, machine, int64(*port))

	var collabSource collabNode = collabNode{machine: machine}

	if *collaborator != "" {
		//? If machine name given, register with central node and get port no for machine name, connect to port+1
		collabSource = getCollabNodeDetails(queryHostDetails, *collaborator)
	}
	// Start CRDT environment
	startCollabEnvironment(machine, collabSource)

	<-quitChannel
	fmt.Println("Program exiting.")
}
func connectToCentralNode(centralHostDetails string) net.Conn {
	conn, err := net.Dial("tcp", centralHostDetails)
	helper.CheckErr(err)
	return conn
}

func registerWithCentralNode(centralHostDetails string, machine string, port int64) {
	conn := connectToCentralNode(centralHostDetails)
	defer conn.Close()
	collabNodeComms := comms_handler.NewRegisterCommsHandler(conn)

	message := &comms_handler.RegisterMessage{
		Machine:    machine,
		ListenPort: port,
		TalkPort:   port + 1,
	}
	collabNodeComms.Send(message)
	response, err := collabNodeComms.Receive()
	helper.CheckErr(err)
	if !response.GetStatus() {
		fmt.Println(response.GetStatus())
		panic("Error in registration")
	}
}

func getCollabNodeDetails(centralHostDetails string, collaborator string) collabNode {
	conn := connectToCentralNode(centralHostDetails)
	defer conn.Close()
	collabNodeComms := comms_handler.NewQueryCommsHandler(conn)
	msg := &comms_handler.QueryMessage{Machine: collaborator}
	collabNodeComms.Send(msg)
	response, err := collabNodeComms.Receive()
	helper.CheckErr(err)
	if response.GetStatus() {
		collabSource := collabNode{collaborator, response.GetListenPort(), response.GetTalkPort()}
		return collabSource
	} else {
		panic("Invalid collab node requested")
	}
}

/*
*
Starts collab environment by initializing struct and channel
*/
func startCollabEnvironment(machine string, collab collabNode) {
	lwwReg = lww.InitializeLWWRegister(machine, machine, 0, "")
	fmt.Println(lwwReg)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Environment stated")
	go startHttpServer()
	if collab.machine == machine {
		go startSrcListeningPort(machine)
	} else {
		go connectToCollabSource(machine, collab)
	}
	fmt.Println("you can start editing")
	for {
		select {
		case <-quitChannel:
			fmt.Println("Received quit signal, stopping receiver.")
			return
		default:
			scanner.Scan()
			input := scanner.Text()
			lwwReg.UpdateLocalState(input)
			fmt.Println("Current state", lwwReg.GetValue())

			// Send the local updates to other peers
			dataChannel <- input
		}
	}
}
func startSrcListeningPort(machine string) {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(*port))
	helper.CheckErr(err)
	for {
		if conn, err := listener.Accept(); err == nil {
			fmt.Println("Node connected")
			communicateWithPeers(machine, conn)
		}
	}

}
func communicateWithPeers(machine string, conn net.Conn) {
	syncCommsHandler := comms_handler.NewSyncCommsHandler(conn)
	go handleReceive(syncCommsHandler)
	go handleSend(syncCommsHandler, machine)
}

func connectToCollabSource(machine string, collab collabNode) {
	conn, err := net.Dial("tcp", collab.machine+":"+strconv.Itoa(int(collab.listenPort)))
	helper.CheckErr(err)
	communicateWithPeers(machine, conn)

}

func handleReceive(syncCommsHandler *comms_handler.SyncCommsHandler) {
	for {
		select {
		case <-quitChannel:
			syncCommsHandler.Close()
			fmt.Println("Received quit signal, stopping receiver.")
			return
		default:
			data, err := syncCommsHandler.Receive()
			helper.CheckErr(err)
			fmt.Println("Received data:", data)
			// call merge method
			peerState := lwwReg.PopulatePeerState(data.GetState().GetId(), int(data.GetState().GetTimestamp()), data.GetState().GetValue())
			lwwReg.Merge(peerState)
			fmt.Println("State updated, Current state:", lwwReg.GetValue())
		}

	}
}

func handleSend(syncCommsHandler *comms_handler.SyncCommsHandler, machine string) {
	for {
		select {
		case data := <-dataChannel: // Hypothetical Send function
			// Send data over network
			fmt.Println("Sending data:", data)
			state := &comms_handler.State{Id: machine, Timestamp: id, Value: data}
			syncCommsHandler.Send(&comms_handler.SyncMessage{Id: id, State: state})
			id++
		case <-quitChannel:
			syncCommsHandler.Close()
			fmt.Println("Received quit signal, stopping sender.")
			return
		}
	}
}

func startHttpServer() {
	http.HandleFunc("/state", httpRequestHandler)
	fmt.Println("Starting server on port:", constants.HTTP_PORT)
	http.ListenAndServe(":"+strconv.Itoa(constants.HTTP_PORT), nil)
}

func httpRequestHandler(w http.ResponseWriter, r *http.Request) {
	select {
	case <-quitChannel:
		fmt.Println("Received quit signal, stopping server.")
		return
	default:
		switch r.Method {
		case "GET":
			// Retrieve the latest state (e.g., from lwwReg or any other data structure)
			latestState := lwwReg.GetValue()

			// Send the latest state as the response
			fmt.Fprintf(w, "Latest state: %s\n", latestState)
		case "POST":
			// Read the data from the POST request
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form data", http.StatusBadRequest)
				return
			}
			inputData := r.FormValue("data") //UI sends data as a form field named "data"
			//update local state
			lwwReg.UpdateLocalState(inputData)
			fmt.Println("Current state", lwwReg.GetValue())
			//Send data over channel
			dataChannel <- inputData
		default:
			fmt.Fprintf(w, "Only GET and POST methods are supported. You tried: %s\n", r.Method)
		}
	}
}
