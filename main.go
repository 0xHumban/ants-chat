package main

import (
	"fmt"
	"os"
)

func main() {
	printDebug(fmt.Sprintf("Current software version: %d", ClientVersion))
	argWithoutProg := os.Args[1:]
	marmots := make([]*Marmot, ClientNumber)
	// open server
	if len(argWithoutProg) > 0 {
		fmt.Println("Start handling chats")
		go openConnection(ServerPort, marmots)
		handleMenu(marmots)
	} else {
		// its client, connect to server
		Debug = false
		var marmot *Marmot
		// TODO: find server ip with udp broadcast
		go connectToServer(ServerIP, &marmot)
		createApp(&marmot)
	}
}
