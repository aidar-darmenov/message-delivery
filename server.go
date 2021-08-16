package main

import (
	"github.com/aidar-darmenov/message-delivery/config"
	"log"
	"net"
	"os"
	"strconv"
)
import "fmt"
import "bufio"
import "strings" // требуется только ниже для обработки примера

func main() {

	fmt.Println("Launching server...")

	cfg := config.NewConfiguration("config/config.json")

	startServer(cfg)
	for i := 0; i < 10; i++ {
		go startClient(cfg)
	}

}

func startClient(cfg *config.Configuration) {
	// Устанавливаем прослушивание порта
	ln, err := net.Listen(cfg.Type, ":"+strconv.Itoa(cfg.Port))
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	if conn != nil {
		// Будем прослушивать все сообщения разделенные \n
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// Распечатываем полученое сообщение
		fmt.Print("Message Received:", string(message))
		// Процесс выборки для полученной строки
		newmessage := strings.ToUpper(message)
		// Отправить новую строку обратно клиенту
		conn.Write([]byte(newmessage + "\n"))
	}
}

func startServer(cfg *config.Configuration) {
	// Подключаемся к сокету
	conn, _ := net.Dial(cfg.Type, cfg.Host+":"+strconv.Itoa(cfg.Port))
	for {
		// Чтение входных данных от stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// Отправляем в socket
		fmt.Fprintf(conn, text+"\n")
		// Прослушиваем ответ
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
