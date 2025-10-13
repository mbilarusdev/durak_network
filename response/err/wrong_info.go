package net_err

type WrongInfoErr struct {
	CustomCode int    `json:"custom_code" example:"14"`
	Msg        string `json:"msg"         example:"Неправильные данные"`
}

func NewWrongInfoErr() *WrongInfoErr {
	return &WrongInfoErr{CustomCode: 14, Msg: "Неправильные данные"}
}
