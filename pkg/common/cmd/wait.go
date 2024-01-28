package cmd

import (
	"log"
	"net"
	"time"
)

func WaitForService(addr string) {
	log.Printf("Waiting for service %s...", addr)
	for {
		log.Printf("Trying to connect to %s...", addr)
		conn, err := net.Dial("tcp", addr)
		if err == nil {
			conn.Close()
			log.Printf("Service %s is up!", addr)
			return
		}
		time.Sleep(time.Millisecond * 500)
	}
}