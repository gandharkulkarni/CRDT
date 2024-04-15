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
	"os"
	"strconv"
	"strings"
)

type collabNode struct {
	machine    string
	listenPort int32
	talkPort   int32
}

func main() {
	port := flag.Int("port", constants.COLLAB_PORT, "Listner port")
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
	fmt.Println("Sender port :", *port+1)
	machine, err := os.Hostname() //machine.domain
	helper.CheckErr(err)
	machine = strings.Split(machine, ".")[0]

	centralHostDetails := *central
	queryPort, err := strconv.Atoi(strings.Split(centralHostDetails, ":")[1])
	helper.CheckErr(err)
	queryHostDetails := strings.Split(centralHostDetails, ":")[0] + ":" + strconv.Itoa(queryPort+1)
	// If machine name not given, register with central node port connect to
	registerWithCentralNode(centralHostDetails, machine, int64(*port))
	if *collaborator != "" {
		//? If machine name given, register with central node and get port no for machine name, connect to port+1
		collabSource := getCollabNodeDetails(queryHostDetails, *collaborator)
	} else {
		// Start CRDT environment
		startCollabEnvironment(machine)
	}
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

func startCollabEnvironment(machine string) {
	var lwwReg *lww.LWWRegister = lww.InitializeLWWRegister(machine, machine, 0, "")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		input := scanner.Text()

		//TODO: Need to write a method to change the state text

		//TODO: listen for updates from collab node
		//TODO: Call merge function

		//TODO: May have to change proto for collab msg

	}
}
