package service

import "sync"

func (s *Service) GetConnectedClientsIds() *sync.Map {
	return s.Clients.Map
}
