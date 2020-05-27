package lib

type Response struct {
	Status  bool    `json:"status"`
	Message interface{}  `json:"message"`
	Data    interface{}  `json:"data"`
}