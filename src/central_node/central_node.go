package main

import (
	"crdt/src/comms_handler"
	"flag"
	"log"
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
	//TODO: create file & class for collaborating nodes

}
func registerCollabNodes(port *int) {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(*port))
	checkErr(err)
	for {
		if conn, err := listener.Accept(); err == nil {
			commsHandler := comms_handler.NewRegisterCommsHandler(conn)
			handleCollabRegistration(commsHandler)
		}
	}
}

func handleCollabRegistration(msgHandler *comms_handler.RegisterCommsHandler) {
	msg, err := msgHandler.Receive()
	checkErr(err)
	_, ok := nodeMap[msg.GetMachine()]
	if !ok {
		nodeMap[msg.GetMachine()] = msg.GetPort()
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}
