package main

import (
	"crdt/src/comms_handler"
	"crdt/src/helper"
	"flag"
	"net"
	"strconv"
)

var nodeMap = map[string]int64{}

func main() {
	port := flag.Int("port", 7000, "Listner port")

	flag.Parse()

	if *port == 0 {
		panic("Insufficient number of arguments. Usage: main.go -port=<port>")
	}

	go registerCollabNodes(port)
	shareCollabNodeDetails(port)

}
func registerCollabNodes(port *int) {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(*port))
	helper.CheckErr(err)
	for {
		if conn, err := listener.Accept(); err == nil {
			commsHandler := comms_handler.NewRegisterCommsHandler(conn)
			handleCollabRegistration(commsHandler)
		}
	}
}

func handleCollabRegistration(msgHandler *comms_handler.RegisterCommsHandler) {
	msg, err := msgHandler.Receive()
	helper.CheckErr(err)
	_, ok := nodeMap[msg.GetMachine()]
	if !ok {
		nodeMap[msg.GetMachine()] = msg.GetPort()
	}
}

func shareCollabNodeDetails(port *int) {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(*port+1))
	helper.CheckErr(err)
	for {
		if conn, err := listener.Accept(); err == nil {
			// TODO: Share details of collab node
		}
	}
}

func handleCollabRequest() {

}
