package main

import (
	"fmt"
	"net"
	"strings"
)

var pairs = make(map[string]*net.TCPConn)

func Server() {
	localAddress := ":8080"
	addr, _ := net.ResolveTCPAddr("tcp", localAddress)
	list, _ := net.ListenTCP("tcp", addr)
	fmt.Println("[INFO] Server started")
	for {
		conn, _ := list.AcceptTCP()
		buffer := make([]byte, 1024)
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			panic(err)
		}

		pairingKey := string(buffer[0:bytesRead])
		fmt.Println("Received key: ", pairingKey)

		if pairs[pairingKey] == nil {
			pairs[pairingKey] = conn
			continue
		} else {
			conn2 := pairs[pairingKey]
			ip1 := strings.Split(conn.RemoteAddr().String(), ":")[0]
			ip2 := strings.Split(conn2.RemoteAddr().String(), ":")[0]
			conn.Write([]byte(ip2))
			conn2.Write([]byte(ip1))
			conn.Close()
			conn2.Close()
			pairs[pairingKey] = nil
			fmt.Printf("[INFO] Key %s paired: %s <-> %s\n", pairingKey, ip1, ip2)
		}
	}
}
