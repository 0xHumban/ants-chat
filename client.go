package main

import (
	"fmt"
	"net"
	"time"
)

var ClientVersion = 5

var UpdateFilePath = "build/client-"
var UpdateFilename = fmt.Sprintf("%s%d", UpdateFilePath, ClientVersion)

const ServerIP = "127.0.0.1:8080"
const RetryDelais = 5

// handle connection client side
// client waiting for server instructions
// 'exit': connection closed
// '1': count 'e' in response
// returns if the connection has been asked by server
func (m *Marmot) handleConnectionClientSide(connectionClosedProperly chan bool) bool {
	defer m.conn.Close()
	for !m.response.isExit() {
		res := m.readResponse()
		if !res {
			printDebug("ERROR reading server response")
			connectionClosedProperly <- false
			return false
		}
		if m.response.isExit() {
			printDebug("EXIT request received")
			connectionClosedProperly <- true
			return true
		}

		m.treatServerResponse()

	}
	connectionClosedProperly <- true
	return true
}

// treats the server response
// choose whats the next step, which function the client have to execute
func (m *Marmot) treatServerResponse() {
	switch m.response.Type {
	case ChatType:
		m.treatChatServerResponse()
	default:
		m.treatStringServerResponse()
	}
}

func (m *Marmot) treatChatServerResponse() {
	// message received from server -> show it
	if m.response.ID == "2" {
		printDebug("New chat received from server")
		chat, err := decodeChat(m.response.Data)
		if err != nil {
			printError(fmt.Sprintf("error during decoding chat data: %s", err))
			return
		}
		// show it
		fmt.Println(chat.String())
	}
}

func (m *Marmot) treatStringServerResponse() {
	// Ping request
	if m.response.ID == "0" {
		printDebug("Ping pong request received")
		m.data = createMessage("0", String, []byte(fmt.Sprintf("'Pong' from @%s", m.conn.LocalAddr().String())))
		_ = m.writeData(true)
		printDebug("Ping pong response sent")

	}
}

// saves the file stored in data on local system
// start it and kill the old one
// TODO: add verification
/*
func connectToServer(ip string) {
	connectionClosedProperly := false

	for !connectionClosedProperly {
		conn, err := net.Dial("tcp", ip)
		if err != nil {
			fmt.Println("ERROR connecting to server", err)
		} else {
			// DEBUG
			printDebug("Local address: " + conn.LocalAddr().String())
			printDebug("Remote address: " + conn.RemoteAddr().String())
			marmot := NewMarmot(conn)
			connectionClosedProperly = marmot.handleConnectionClientSide()
		}
		if !connectionClosedProperly {
			time.Sleep(RetryDelais * time.Second)
		}
	}
}

*/

func connectToServer(ip string) {
	connectionClosedProperly := make(chan bool, 1)
	connectionClosedProperly <- false
	for !<-connectionClosedProperly {
		conn, err := net.Dial("tcp", ip)
		if err != nil {
			fmt.Println("ERROR connecting to server", err)
			connectionClosedProperly <- false
		} else {
			// DEBUG
			printDebug("Local address: " + conn.LocalAddr().String())
			printDebug("Remote address: " + conn.RemoteAddr().String())
			marmot := NewMarmot(conn)
			go marmot.handleConnectionClientSide(connectionClosedProperly)
		}
		if !<-connectionClosedProperly {
			time.Sleep(RetryDelais * time.Second)
			connectionClosedProperly <- false
		} else {
			connectionClosedProperly <- true
		}
	}
}
