package main

import (
	"crdt/src/helper"
	"flag"
	"os"
	"strings"
)

func main() {
	port := flag.Int("port", 7000, "Listner port")
	central := flag.String("central", "", "Central machine:port")
	collaborator := flag.String("collab", "", "Other collaborator node name")

	flag.Parse()

	if *port == 0 {
		panic("Insufficient number of arguments. Usage: main.go -port=<port>")
	}
	if *central == "" {
		panic("Insufficient number of arguments. Usage: main.go -central=<machine:port>")
	}

	machine, err := os.Hostname()
	helper.CheckErr(err)

	details := strings.Split(*central, ":")
	centralNodeName := details[0]
	centralNodePort := details[1]

	// TODO: If machine name not given, register with central node port connect to
	registerWithCentralNode()
	if *collaborator != "" {
		// TODO: If machine name given, register with central node and get port no for machine name, connect to port+1
		getCollabNodeDetails()
	}

}

func registerWithCentralNode() {

}

func getCollabNodeDetails() {

}
