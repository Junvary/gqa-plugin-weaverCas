package model

type ValidateTicket struct {
	AppId   string `json:"app_id"`
	Ticket  string `json:"ticket"`
	Service string `json:"service"`
}
