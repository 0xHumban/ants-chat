package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Represents Chat between client
// Chat:
//      Client name
//      Text
// maybe add ip address to identify machine

type Chat struct {
	FromClientName string
	Text           string
}

// format the message to nicely view in the terminal
func (c Chat) String() string {
	return fmt.Sprintf("%s | %s", c.FromClientName, c.Text)
}

func (c Chat) encode() ([]byte, error) {
	return encode(c)
}

// deserializes the struct
func decodeChat(data []byte) (*Chat, error) {
	return decode[Chat](data)
}

// sends messages to all Ants connected to server except the from one
func (ms Marmots) publishChat(c Chat) {
	// NOTE: Maybe do not Pings all clients to gain latency between messages
	// NOTE 2: BUG if Pings before: error decoding data
	// ms.Pings()
	printDebug(fmt.Sprintf("\nPublish Chat start, Chat: '%s'", c.String()))
	// add chat to all marmot except the from one
	chatEncoded, err := c.encode()
	if err != nil {
		printError(fmt.Sprintf("During publish chat, encoding chat: %s", err))
		return
	}
	for _, m := range ms {
		if m != nil {
			m.data = createMessage("2", ChatType, chatEncoded)
		}
	}
	ms.performAction((*Marmot).publishChat)
	// if a client disconnect, remove it
	for i, m := range ms {
		if m != nil && !<-m.end {
			ms[i] = nil
			printDebug("@" + m.conn.RemoteAddr().String() + " has been removed of the clients list")
		}
	}
	printDebug(fmt.Sprintf("Publish Chat end, Chat: '%s'", c.String()))
}

// send Chat to client
func (m *Marmot) publishChat() {
	m.SendData("publish chat", true)
}

func (m *Marmot) SendChat(chat Chat) {
	printDebug("Start send chat to server")
	chatEncoded, err := chat.encode()
	if err != nil {
		printError(fmt.Sprintf("During sending chat, encoding chat: %s", err))
		return
	}
	m.data = createMessage("2", ChatType, chatEncoded)
	m.start <- true
	m.SendData("send chat to server", true)
	<-m.end
	printDebug("End send chat to server")
}

// handle all chats received and publish them to others client
func (ms Marmots) handleServerChats() {
	fmt.Println("Start handling chats")
	for _, m := range ms {
		if m != nil {
			go m.handleServerChatsForOne(ms)
		}
	}

}

func (m *Marmot) handleServerChatsForOne(ms Marmots) {
	// TODO: add something to stop the loop
	for {

		res := m.readResponse()
		if !res {
			printDebug("ERROR reading client response")
			m.Close()
			printDebug("@" + m.conn.RemoteAddr().String() + " has been removed of the clients list")
			m = nil
			break
		}

		printDebug(fmt.Sprintf("New chat received from client: @%s", m.conn.RemoteAddr()))
		chat, err := decodeChat(m.response.Data)
		if err != nil {
			printError(fmt.Sprintf("error during decoding chat data: %s", err))
			m.Close()
			printDebug("@" + m.conn.RemoteAddr().String() + " has been removed of the clients list")
			m = nil
			break
		}
		fmt.Println(chat.String())
		ms.publishChat(*chat)

	}

}

// handle menu for server admin to send a test chat to clients connected
func handlePublishChatTestMenu(ms Marmots) {
	scanner := bufio.NewScanner(os.Stdin)
	c := Chat{"SERVER", ""}
	for {
		fmt.Println("======= Send TEST CHAT TO CLIENTS ======= ")
		fmt.Println("It will send a test chat to all clients connected")
		fmt.Println("Enter the message you want to send")
		fmt.Println("(Enter '-1' to leave)")

		scanner.Scan()

		choice := strings.TrimSpace(scanner.Text())
		if choice == "-1" {
			return
		} else {
			c.Text = choice
			fmt.Println(c.String())
			ms.publishChat(c)
			return
		}
	}
}
