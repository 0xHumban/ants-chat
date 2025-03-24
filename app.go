package main

import (
	// "fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var ChatContainer *fyne.Container = nil

func createApp(server **Marmot) {

	myApp := app.New()

	myWindow := myApp.NewWindow("Message App")

	// Créer un conteneur pour les messages
	ChatContainer = container.NewVBox()

	// Créer un champ de saisie
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Entrez votre message...")

	// Créer un bouton pour envoyer le message
	sendButton := widget.NewButton("Envoyer", func() {
		chat := Chat{(*server).Name, entry.Text}
		(*server).SendChat(chat)
		// addMessage(entry.Text)
		entry.SetText("")
	})

	// Gérer l'appui sur la touche Entrée dans le champ de saisie
	entry.OnSubmitted = func(text string) {
		chat := Chat{(*server).Name, text}
		(*server).SendChat(chat)
		// addMessage(text)
		entry.SetText("")
	}

	// Créer un conteneur pour le champ de saisie et le bouton
	inputContainer := container.NewVBox(entry, sendButton)

	// Créer un conteneur principal avec la zone de messages en haut et le champ de saisie en bas
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
