package net_err

type UnmarshalingErr struct {
	CustomCode int    `json:"custom_code" example:"11"`
	Msg        string `json:"msg"         example:"Ошибка декодирования тела запроса"`
}

func NewUnmarshalingErr() *UnmarshalingErr {
	return &UnmarshalingErr{CustomCode: 11, Msg: "Ошибка декодирования тела запроса"}
}
