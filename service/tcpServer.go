package service

import (
	"bufio"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net"
	"strconv"
	"sync"
)

func (s *Service) StartTcpServer() {
	l, err := net.Listen(s.GetConfigParams().ListenerType, s.GetConfigParams().ListenerHost+":"+strconv.Itoa(s.GetConfigParams().ListenerPort))
	if err != nil {
		return
	}

	defer l.Close()

	// Using sync.Map to store connected clients
	var connMap = &sync.Map{}

	for {
		conn, err := l.Accept()
		if err != nil {
			s.Logger.Error("error accepting connection", zap.Error(err))
			return
		}

		id := uuid.New().String()
		connMap.Store(id, conn)

		go s.HandleUserConnection(id, conn)
	}
}

func (s *Service) HandleUserConnection(id string, c net.Conn) {
	defer func() {
		c.Close()
		s.Clients.Delete(id)
	}()

	for {
		userInput, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			s.Logger.Error("error reading from client", zap.Error(err))
			return
		}

		s.Clients.Range(func(key interface{}, value interface{}) bool {
			if conn, ok := value.(net.Conn); ok {
				if _, err := conn.Write([]byte(userInput)); err != nil {
					s.Logger.Error("error on writing to connection", zap.Error(err))
				}
			}
			return true
		})
	}
}
