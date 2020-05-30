package main

import (
	"flag"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	for {
		l := make([]byte, 1024)
		s, err := conn.Read(l)
		if err != nil {
			log.Printf("Error reading from connection: %s", err)
			conn.Close()
			return
		}
		if verbose {
			log.Printf("Data received: %s", l[:s])
		}
		conn.Write(l[:s])
	}
}

var verbose bool

func main() {

	// setup command line params
	var port string
	var ip string
	flag.StringVar(&port, "p", "0042", "Port to listen on (default: 0042)")
	flag.StringVar(&ip, "i", "0.0.0.0", "IP to listen on (default: 0.0.0.0)")
	flag.BoolVar(&verbose, "v", false, "Enable verbose logging")
	flag.Parse()

	if verbose {
		log.Printf("Attempting to listen on %s:%s", ip, port)
	}

	// start listener
	address := ip + ":" + port
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Could not bind to %s: %s", address, err)
	}

	// kick off new go routines for each connection
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
			continue
		}
		go handleConnection(conn)
	}
}
