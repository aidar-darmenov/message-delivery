package service

import (
	"bufio"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net"
	"strconv"
)

func (s *Service) StartTcpServer() {
	l, err := net.Listen(s.GetConfigParams().ListenerType, s.GetConfigParams().ListenerHost+":"+strconv.Itoa(s.GetConfigParams().ListenerPort))
	if err != nil {
		return
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			s.Logger.Error("error accepting connection", zap.Error(err))
			return
		}

		id := uuid.New().String()
		s.Clients.Map.Store(id, conn)
		s.Clients.Ids = append(s.Clients.Ids, id)

		go s.HandleUserConnection(id, conn)
	}
}

func (s *Service) HandleUserConnection(id string, c net.Conn) {
	defer func() {
		c.Close()
		s.Clients.Map.Delete(id)
	}()

	for {
		userInput, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			s.Logger.Error("error reading from client", zap.Error(err))
			return
		}

		s.Clients.Map.Range(func(key interface{}, value interface{}) bool {
			if conn, ok := value.(net.Conn); ok {
				if err := s.SendMessageToClient(conn, userInput); err != nil {
					s.Logger.Error("error on writing to connection", zap.Error(err))
				}
			}
			return true
		})
	}
}
