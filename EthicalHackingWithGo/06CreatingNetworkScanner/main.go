package main

import (
	"fmt"
	"net"
)

func main() {
	IP := "scanme.nmap.org"
	//	Port := "80"

	for i := 1; i < 100; i++ {
		address := fmt.Sprintf(IP+":%d", i)
		connection, err := net.Dial("tcp", address)
		if err == nil {
			fmt.Printf("[+] Connection established..PORT %v %s\n", i, connection.RemoteAddr().String())
		}
	}
}
