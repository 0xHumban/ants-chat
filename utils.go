package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const RedColor = "\033[31m"
const YellowColor = "\033[33m"
const ResetColor = "\033[0m"
const CyanColor = "\033[36m"

var Debug = true
var HighDebug = false

func printDebugCondition(text string, show bool) {
	if show {
		printDebug(text)
	}
}

func printDebug(text string) {
	if Debug {
		now := time.Now()
		millis := fmt.Sprintf("%d", now.UnixMilli())
		fmt.Println(YellowColor + millis + "| DEBUG: " + text + ResetColor)
	}
}

func printHighDebug(text string) {
	if HighDebug {
		now := time.Now()
		millis := fmt.Sprintf("%d", now.UnixMilli())
		fmt.Println(CyanColor + millis + "| HIGH DEBUG: " + text + ResetColor)
	}
}

func printError(text string) {
	now := time.Now()
	millis := fmt.Sprintf("%d", now.UnixMilli())
	fmt.Println(RedColor + millis + "| ERROR: " + text + ResetColor)
}

func showMenu() {
	fmt.Println("\n===== Menu ===== ")
	fmt.Println("3. Close connections")
	fmt.Println("4. Send test Chat to all clients")
	fmt.Println("6. Exit (will let clients trying to reconnect to server)")
	fmt.Print("Choose an option:\n")
}

func handleMenu(marmots Marmots) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMenu()
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "3":
			marmots.CloseConnections()
		case "4":
			handlePublishChatTestMenu(marmots)
		case "6":
			return
		default:
			printError("Invalid option, please try again.")
		}
	}
}
