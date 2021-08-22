package service

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/aidar-darmenov/message-delivery/model"
	"go.uber.org/zap"
	"net"
)

func (s *Service) SendMessageToClientsByIds(message model.MessageToClients) *model.Exception {
	for i := range message.Ids {
		m, ok := s.Clients.Map.Load(message.Ids[i])
		if !ok {
			s.Logger.Error("Failed to load client by ID: " + message.Ids[i])
		}
		if conn, ok := m.(*net.TCPConn); ok {
			if err := s.SendMessageToClient(conn, message.Text); err != nil {
				s.Logger.Error("error on writing to connection", zap.Error(err))
			}
		}
	}
	return nil
}

func (s *Service) SendMessageToClient(conn *net.TCPConn, message string) error {

	var (
		data         = []byte(message)
		msg_len_data = make([]byte, 2)
		buf          = bytes.Buffer{}
	)

	binary.BigEndian.PutUint16(msg_len_data, uint16(len(data)))

	fmt.Println("content length bytes: ", msg_len_data)
	fmt.Println("content bytes: ", data)

	buf.Write(msg_len_data)
	buf.Write(data)

	_, err := conn.Write(buf.Bytes())
	if err != nil {
		return err
	}

	buf.Reset()

	return nil
}
