package service

import (
	"bufio"
	"github.com/aidar-darmenov/message-delivery/helpers"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"log"
	"net"
	"strconv"
)

func (s *Service) StartTcpServer() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", net.JoinHostPort(s.Configuration.Params().ListenerHost, strconv.Itoa(s.Configuration.Params().ListenerPort)))
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.ListenTCP(s.GetConfigParams().ListenerType, tcpAddr)
	if err != nil {
		return
	}

	defer l.Close()

	for {
		conn, err := l.AcceptTCP()
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

func (s *Service) HandleUserConnection(id string, c *net.TCPConn) {
	defer func() {
		c.Close()
		s.Clients.Map.Delete(id)
		helpers.DeleteIdFromClientList(s.Clients, id)
	}()

	for {
		userInput, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			s.Logger.Error("error reading from client", zap.Error(err))
			return
		}

		s.Clients.Map.Range(func(key interface{}, value interface{}) bool {
			if conn, ok := value.(*net.TCPConn); ok {
				if err := s.SendMessageToClient(conn, userInput); err != nil {
					s.Logger.Error("error on writing to connection", zap.Error(err))
				}
			}
			return true
		})
	}
}
