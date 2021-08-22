package model

type MessageToClients struct {
	Ids  []string `json:"ids"`
	Text string   `json:"text"`
}
