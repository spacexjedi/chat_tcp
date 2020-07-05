package main

import (
  "fmt"
  "log"
  "net"
)

func main() {
	fmt.Println("Hello World")

	s:= newServer()
	go s.run()

	listener, err := net.Listen("tcp", ":8888")

	for {

		conn, err := listener.Accept()
		if err != nill { // diferente
			log.Printf("failed to accept connection %s", err.Error())
			continue
		}

		go s.newClient(conn)

	} // end of for
}
