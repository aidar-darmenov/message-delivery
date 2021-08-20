package model

import "sync"

type Clients struct {
	Map *sync.Map
	Ids []string
}
