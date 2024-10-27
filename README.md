## TCP/UDP server/client

### Description
This project is a simple TCP/UDP server/client implementation

Client sends a message to the server and 
the server responds with the message about bytes received

Need run the server first and then the client
If client is disconnected, the server can connect to another client

### How to run
1. Run the TCP server
```./start -mode=tcp-server -ip=[ip]```
2. Run the TCP client
```./start -mode=tcp-client -ip=[ip]```
3. Run the UDP client
```./start -mode=udp-client -ip=[ip]```
4. Run the UDP server
```./start -mode=udp-server -ip=[ip]```