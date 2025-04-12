package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var ChatContainer *fyne.Container = nil

func createApp(server **Marmot) {

	myApp := app.New()

	myWindow := myApp.NewWindow("Message App")

	showInputDialog(myWindow, server)

	ChatContainer = container.NewVBox()

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Entrez votre message...")

	sendButton := widget.NewButton("Envoyer", func() {
		chat := Chat{(*server).Name, entry.Text}
		(*server).SendChat(chat)
		entry.SetText("")
	})

	entry.OnSubmitted = func(text string) {
		chat := Chat{(*server).Name, text}
		(*server).SendChat(chat)
		entry.SetText("")
	}

	inputContainer := container.NewVBox(entry, sendButton)

	content := container.NewBorder(nil, inputContainer, nil, nil, container.NewVScroll(ChatContainer))

	myWindow.SetContent(content)

	myWindow.Resize(fyne.NewSize(400, 600))

	myWindow.ShowAndRun()
}

// Fonction pour ajouter un message à la zone de visualisation
func addMessage(text string) {
	if text != "" {
		ChatContainer.Add(widget.NewLabel(text))
		ChatContainer.Refresh()
	}
}

// Function to show the input dialog and retrieve user input
func showInputDialog(window fyne.Window, server **Marmot) {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter your name...")

	ipEntry := widget.NewEntry()
	ipEntry.SetText(ServerIP) // Default IP address

	// Function to check inputs and re-open dialog if empty
	var showDialog func()
	showDialog = func() {
		form := container.NewVBox(
			widget.NewLabel("Please enter your information:"),
			widget.NewLabel("Name:"),
			nameEntry,
			widget.NewLabel("Server IP:"),
			ipEntry,
		)

		dialog.ShowCustomConfirm("Required Information", "Confirm", "", form, func(confirmed bool) {
			if nameEntry.Text == "" {
				showDialog() // Reopen if name is empty (IP has a default)
			} else {
				fmt.Println("Name:", nameEntry.Text)
				fmt.Println("Server IP:", ipEntry.Text)
				ServerIP = ipEntry.Text
				ClientName = nameEntry.Text
				go connectToServer(ServerIP, server)
				window.Show()
			}
		}, window)
	}

	showDialog() // Show dialog initially

}
