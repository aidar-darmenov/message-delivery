package service

import (
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
		if conn, ok := m.(net.Conn); ok {
			if err := s.SendMessageToClient(conn, message.Text); err != nil {
				s.Logger.Error("error on writing to connection", zap.Error(err))
			}
		}
	}
	return nil
}

func (s *Service) SendMessageToClient(conn net.Conn, message string) error {
	_, err := conn.Write([]byte(message))
	if err != nil {
		return err
	}
	return nil
}
