package main

import (
	"fmt"
	"net"
)

var pairs = make(map[string]*net.TCPConn)

func Server() {
	localAddress := ":8080"
	addr, _ := net.ResolveTCPAddr("tcp", localAddress)
	list, _ := net.ListenTCP("TCP", addr)
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

		ip1 := conn.RemoteAddr().String()
		ip2 := ""
		if pairs[pairingKey] == nil {
			pairs[pairingKey] = conn
			continue
		} else {
			conn2 := pairs[pairingKey]
			conn.Write([]byte(conn2.RemoteAddr().String()))
			conn2.Write([]byte(conn.RemoteAddr().String()))
			conn.Close()
			conn2.Close()
			pairs[pairingKey] = nil
			fmt.Printf("[INFO] Key %s paired: %s <-> %s\n", pairingKey, ip1, ip2)
		}

		//r1, _ := net.ResolveUDPAddr("udp", ip1)
		//r2, _ := net.ResolveUDPAddr("udp", ip2)
		//conn.Write([]byte(ip1))
		//conn.WriteTo([]byte(ip2), r1)
		//conn.WriteTo([]byte(ip1), r2)

	}
}
