package main

import (
	"bufio"
	"fmt"
	"github.com/aidar-darmenov/message-delivery/config"
	"github.com/prometheus/common/log"
	"net"
	"strconv"
	"strings"
)

func main() {

	cfg := config.NewConfiguration("config/config.json")

	StartServer(cfg)
}

func StartServer(cfg *config.Configuration) {

	log.Info("Launching server...")

	ln, err := net.Listen(cfg.Type, ":"+strconv.Itoa(cfg.Port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil && conn != nil {
			//err handling
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed. Error: ", err)
			conn.Close()
			return
		}
		fmt.Print("Message Received:", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}
}
