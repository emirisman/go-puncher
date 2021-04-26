package main

import (
	"fmt"
	"net"
)

var pairs = make(map[string]string)

// Server --
func Server() {
	localAddress := ":8080"

	addr, _ := net.ResolveUDPAddr("udp", localAddress)
	conn, _ := net.ListenUDP("udp", addr)
	fmt.Println("[INFO] Server started")
	for {
		buffer := make([]byte, 1024)
		bytesRead, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}

		pairingKey := string(buffer[0:bytesRead])
		fmt.Println("Received key: ", pairingKey)

		ip1 := remoteAddr.String()
		ip2 := ""
		if pairs[pairingKey] == "" {
			pairs[pairingKey] = ip1
			continue
		} else {
			ip2 = pairs[pairingKey]
		}

		r1, _ := net.ResolveUDPAddr("udp", ip1)
		r2, _ := net.ResolveUDPAddr("udp", ip2)
		conn.WriteTo([]byte(ip2), r1)
		conn.WriteTo([]byte(ip1), r2)
		fmt.Printf("[INFO] Key %s paired: %s <-> %s\n", pairingKey, ip1, ip2)

	}
}
