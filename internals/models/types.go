package models

type Event struct {
	Operation    string `json:"op"`
	Entity       string `json:"e"`
	ID           int    `json:"id"`
	Member       string `json:"m"`
	Microservice string `json:"ms"`
}

type Response struct {
	MessageID *string
}
