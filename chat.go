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
			ms.publishChat(c)
			return
		}
	}
}
