package main

import (
	"fmt"
	"net"
	"os"
)
// main function
func main() {
	// create localhost tcp listener
	listener, err := net.Listen("tcp","127.0.0.1:8888")
	ServerPrintErr(err, "net listen")
	
	// for each client, go a thread
	for {
		conn,e := listener.Accept()
		ServerPrintErr(e,"listener.accept")
		go ChatWith(conn)
	}
}

// print err func
func ServerPrintErr(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}

// function for a single client
func ChatWith(conn net.Conn){
	buff := make([]byte, 1024)
	for {
		// read msg from client
		n,err := conn.Read(buff)
		ServerPrintErr(err,"conn.read buffer")
		//change msg to string type
		clientMsg := string(buff[0:n])
		fmt.Printf("recv msg",conn.RemoteAddr(),clientMsg)
		// response
		if clientMsg != "im off" {
			conn.Write([]byte("server read:"+clientMsg))
		} else {
			conn.Write([]byte("bye"))
            break
		}
	}
	// close connection
	conn.Close()
	fmt.Printf("server disable connection",conn.RemoteAddr())
}