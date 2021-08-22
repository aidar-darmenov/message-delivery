package model

import "sync"

type Clients struct {
	Map    *sync.Map
	Params []ClientParams
}

type ClientParams struct {
	Id       string `json:"-"`
	HttpPort int    `json:"http_port"`
	Name     string `json:"name"`
}
