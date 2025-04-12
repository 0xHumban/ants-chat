package main

// Represents the chat client (Ant)
// with a name and a Body function (Marmot)
type Ant struct {
	Name   string
	Marmot Marmot
}

// represents all clients connected to server
type Ants []*Ant

// sends messages to all Ants connected to server except the from one
