## Notes related to this project

## Step 1: Client can send messages to server + receive messages

#### Message ID
Message received from server and show it: `2`

## Step 2: Server broadcast to others clients new messages received

## Step 3: Network discovery
At the start, the client don't know the server IP, to get it, client will use udp broadcast.

### Windows export
```bash
#GOOS=windows GOARCH=amd64 fyne package -os windows -icon ./assets/ants.png
fyne package -os windows -name "AntsChat.exe" -icon ./assets/logo.png
GOOS=windows GOARCH=amd64 go build -o ants_chat_v9.exe



```

