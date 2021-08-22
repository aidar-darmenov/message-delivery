package helpers

import (
	"github.com/aidar-darmenov/message-delivery/model"
)

func DeleteIdFromClientList(clients *model.Clients, id string) bool {

	for i := range clients.Ids {
		if clients.Ids[i] == id {
			clients.Ids = append(clients.Ids[:i], clients.Ids[i+1:]...)
		}
	}

	return true
}
