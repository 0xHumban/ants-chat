# ants-chat
A Local Network chat built in Go

The goal is to build a local network chat where clients can connect to server, send / receive message to other clients.


### Network discovery
At the start, the client don't know the server IP, to get it, client will use udp broadcast.
