package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)


func main() {
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	ClientPrintErr(err, "client conn error")

	buffer := make([]byte, 1024)
	reader := bufio.NewReader(os.Stdin)

	for {
		lineBytes,_,_ := reader.ReadLine()
		conn.Write(lineBytes)
		n,err := conn.Read(buffer)
		ClientPrintErr(err, "client read error")

		serverMsg := string(buffer[0:n])
        fmt.Printf("server msg",serverMsg)
        if serverMsg == "bye" {
            break
        }
	}
}

// print err func
func ClientPrintErr(err error, when string) {
	if err != nil {
		fmt.Println(err, when)
		os.Exit(1)
	}
}