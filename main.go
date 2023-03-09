package main

import (
	"fmt"
	"net"
)

var domain string = "google.com"

func main() {
	// Resolve the domain name using the default DNS resolver
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("Error resolving domain name:", err)
		return
	}

	// Print the IP addresses associated with the domain name
	for _, ip := range ips {
		fmt.Println(domain, "->", ip)
	}
}
