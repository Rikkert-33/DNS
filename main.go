package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Print("Name a website: ")
	reader := bufio.NewReader(os.Stdin)
	domain, _ := reader.ReadString('\n')
	domain = strings.TrimSpace(domain)
	// Resolve the domain name using the default DNS resolver
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("Error resolving domain name:", err)
		return
	}

	// Print the IP addresses with the domain name
	for _, ip := range ips {
		fmt.Println(domain, "->", ip)
	}
}
