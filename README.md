# ants-chat


🐜 - A Local Network chat built in Go

The goal is to build a local network chat where clients can connect to server, send / receive message to other clients.
This project is designed to be educational and fun, providing a hands-on experience with distributed systems and concurrent programming in Go.

## 🎬 Demo Video

![ANts chat demo](assets/demo.gif)
---

![Ants chat schema](assets/ants-chat-schem.png)

--- 

##  Features

### Server

- **handle client connection**, receiving / sending message to connected clients
- **send message to clients**, send message has 'SERVER' to all clients
- **Close client remotely**, can close all current connections


### Client

- **UI interface**, views / sends message to the server.
- **Send message to sever**, and server returns to others clients. 
- **Persistent connection**:
  - If the connection is lost, the client automatically attempts to reconnect.
  - It continues until the server sends an exit message.


### Cross-platform support 

You can send messages **from Windows or Linux** computers on the same network.


### Running the Project


1. Start the server:
   
   ```sh
   ./ants-chat server
   ```

2. Start one or more clients:
   
   ```sh
   ./ants-chat
   ```

---

## How It Works

1. **Start the server**: it waits for client connections.
2. **Clients connect**: each client registers with the server.
3. **The server handle client messgaes**: receive message and returns to others clients.
4. **If a client disconnects**, it automatically attempts to reconnect.




---

This project is inspired from another one [Marmot Reduce](https://github.com/0xHumban/marmot-reduce)



---
## License

This project is open-source and licensed under **MIT**.

