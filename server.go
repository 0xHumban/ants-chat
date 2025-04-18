package main

import (
	"fmt"
	"net"
)

const ServerPort = ":8080"
const ClientNumber = 20
const TimeoutServerRequestSeconds = 20000000

// open a port to allow client to connect
// In:
// - port: port to open
// - handleFct: function pointer to handle different connexions
func openConnection(port string, marmots Marmots) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("ERROR during listening:", err)
	}

	// get network interfaces
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		printError(fmt.Sprintf("Error getting network interfaces:", err))
		return
	}

	printDebug("Server is listening on:")
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			printDebug(fmt.Sprintf(" - %s%s\n", ipNet.IP.String(), port))
		}
	}

	defer ln.Close()

	printDebug("Server waiting for connections")
	for marmots.clientsLen() < ClientNumber {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ERROR accepting connection: ", err)
			continue
		} else {
			printDebug("New client connected: @" + conn.RemoteAddr().String())
		}
		// search next marmot index to insert
		for i := 0; i < ClientNumber; i++ {
			if marmots[i] == nil {
				marmots[i] = NewMarmot(conn)
				go marmots[i].handleServerChatsForOne(marmots)
				break
			}
		}

	}

}
