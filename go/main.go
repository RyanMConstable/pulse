package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, ifi := range interfaces {
		fmt.Printf("Name: %v\n", ifi.Name)
		fmt.Printf("MAC Address: %v\n", ifi.HardwareAddr)
		fmt.Printf("MTU: %v\n", ifi.MTU)
		fmt.Printf("Flags: %v\n", ifi.Flags)
	}
}
