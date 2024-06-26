package main

import (
	"crdt/src/comms_handler"
	"crdt/src/constants"
	"crdt/src/helper"
	"flag"
	"fmt"
	"net"
	"strconv"
)

type port struct {
	listenPort int32
	talkPort   int32
}

var nodeMap = make(map[string]port)

func main() {
	port := flag.Int("port", constants.CENTRAL_PORT, "Listner port")

	flag.Parse()

	if *port == 0 {
		panic("Insufficient number of arguments. Usage: main.go -port=<port>")
	}

	//* Register collab nodes
	registrationPort := strconv.Itoa(*port)
	go registerCollabNodes(registrationPort)

	//* Share details of registered nodes
	detailDispatcherPort := strconv.Itoa(*port + 1)
	shareCollabNodeDetails(detailDispatcherPort)

}
func registerCollabNodes(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	helper.CheckErr(err)
	fmt.Println("Registering collab nodes on port: ", port)
	for {
		if conn, err := listener.Accept(); err == nil {
			fmt.Println("Node connected")
			commsHandler := comms_handler.NewRegisterCommsHandler(conn)
			handleCollabRegistration(commsHandler)
		}
	}
}

func handleCollabRegistration(msgHandler *comms_handler.RegisterCommsHandler) {
	msg, err := msgHandler.Receive()
	helper.CheckErr(err)

	fmt.Println("Message received from ", msg.GetMachine())
	_, ok := nodeMap[msg.GetMachine()]
	if !ok {
		nodeMap[msg.GetMachine()] = port{listenPort: int32(msg.GetListenPort()), talkPort: int32(msg.GetTalkPort())}
	}
	msgHandler.Send(&comms_handler.RegisterMessage{Machine: msg.GetMachine(), ListenPort: msg.GetListenPort(), TalkPort: msg.GetTalkPort(), Status: true})
}

func shareCollabNodeDetails(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	helper.CheckErr(err)
	fmt.Println("Listening for collab nodes on port: ", port)
	for {
		if conn, err := listener.Accept(); err == nil {
			// Share details of collab node
			queryCommsHandler := comms_handler.NewQueryCommsHandler(conn)
			handleCollabRequest(queryCommsHandler)
		}
	}
}

func handleCollabRequest(queryCommsHandler *comms_handler.QueryCommsHandler) {
	queryMsg, err := queryCommsHandler.Receive()
	helper.CheckErr(err)
	machine := queryMsg.GetMachine()
	ports, ok := nodeMap[machine]
	if ok {
		response := &comms_handler.QueryMessage{
			Machine:    machine,
			ListenPort: ports.listenPort,
			TalkPort:   ports.talkPort,
			Status:     true,
		}
		queryCommsHandler.Send(response)
	} else {
		// send invalid response
		queryCommsHandler.Send(&comms_handler.QueryMessage{Status: false})
	}
}
