package net_err

type ServerErr struct {
	CustomCode int    `json:"custom_code" example:"1"`
	Msg        string `json:"msg"         example:"Ошибка сервера"`
}

func NewServerErr() *ServerErr {
	return &ServerErr{CustomCode: 1, Msg: "Ошибка сервера"}
}
