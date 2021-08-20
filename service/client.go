package service

import "sync"

func (s *Service) GetConnectedClients() *sync.Map {
	return s.Clients
}
