package service

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/aidar-darmenov/message-delivery/model"
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

	go s.HandleOutgoingTraffic()

	for {
		conn, err := l.AcceptTCP()
		if err != nil {
			s.Logger.Error("error accepting connection", zap.Error(err))
			return
		}

		id := uuid.New().String()
		s.Clients.Map.Store(id, conn)
		s.Clients.Params = append(s.Clients.Params, model.ClientParams{
			Id: id,
		})

		go s.HandleClients(id, conn)
	}
}

func (s *Service) HandleClients(id string, conn *net.TCPConn) {
	defer func() {
		conn.Close()
		s.Clients.Map.Delete(id)
		s.DeleteIdFromClientList(id)
	}()

	s.GetClientParams(conn, id)

	s.HandleIncomingTraffic(conn)
}

func (s *Service) HandleOutgoingTraffic() {
	for {
		select {
		case message := <-s.ChannelMessages:
			s.SendMessageToClientsByIds(message)
		}
	}
}

func (s *Service) GetClientParams(conn *net.TCPConn, id string) {
	var buf [2048]byte
	var contentLength int

	// Reading content length
	n, err := conn.Read(buf[:2])
	e, ok := err.(net.Error)

	if err != nil && ok && !e.Timeout() {
		s.Logger.Error("Error reading client params length from TCP connection", zap.Error(err))
		return
	}

	if n > 0 {
		contentLength = s.GetContentLength(buf[:n])
	} else {
		conn.Write([]byte("n<0"))
	}

	// Reading content
	n, err = conn.Read(buf[:contentLength])
	e, ok = err.(net.Error)

	if err != nil && ok && !e.Timeout() {
		s.Logger.Error("Error reading client params from TCP connection", zap.Error(err))
		return
	}

	var params model.ClientParams

	if n > 0 {
		params, err = s.ProcessParams(buf[:n])
		if err != nil {
			s.Logger.Error("Error processing client params from TCP connection", zap.Error(err))
		}
	} else {
		conn.Write([]byte("n<0"))
	}

	// Saving params for client
	for i := range s.Clients.Params {
		if s.Clients.Params[i].Id == id {
			s.Clients.Params[i] = model.ClientParams{
				Id:       params.Id,
				HttpPort: params.HttpPort,
				Name:     params.Name,
			}
		}
	}
}

func (s *Service) ProcessParams(buf []byte) (params model.ClientParams, err error) {
	err = json.Unmarshal(buf, &params)
	if err != nil {
		return
	}
	s.Logger.Info(fmt.Sprintf("client params: %v", params))
	return
}

func (s *Service) HandleIncomingTraffic(conn *net.TCPConn) {
	for {
		var buf [2048]byte
		var contentLength int

		// Reading content length
		n, err := conn.Read(buf[:2])
		e, ok := err.(net.Error)

		if err != nil && ok && !e.Timeout() {
			s.Logger.Error("Error reading content length from TCP connection", zap.Error(err))
			break
		}

		if n > 0 {
			contentLength = s.GetContentLength(buf[:n])
		} else {
			conn.Write([]byte("n<0"))
		}

		// Reading content
		n, err = conn.Read(buf[:contentLength])
		e, ok = err.(net.Error)

		if err != nil && ok && !e.Timeout() {
			s.Logger.Error("Error reading content from TCP connection", zap.Error(err))
			break
		}

		var message model.MessageToClients

		if n > 0 {
			message, err = s.ProcessMessage(buf[:n])
			if err != nil {
				s.Logger.Error("Error processing message from TCP connection", zap.Error(err))
			}
		} else {
			conn.Write([]byte("n<0"))
		}

		s.ChannelMessages <- message
	}
}

func (s *Service) GetContentLength(bufContentLength []byte) int {
	cl := int(binary.BigEndian.Uint16(bufContentLength))
	s.Logger.Info(fmt.Sprintf("content length: %d", cl))
	return cl
}

func (s *Service) ProcessMessage(buf []byte) (message model.MessageToClients, err error) {
	err = json.Unmarshal(buf, &message)
	if err != nil {
		return
	}
	s.Logger.Info(fmt.Sprintf("content: %v", message))
	return
}
