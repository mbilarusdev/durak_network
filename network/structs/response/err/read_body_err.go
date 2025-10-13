package net_err

type ReadBodyErr struct {
	CustomCode int    `json:"custom_code" example:"2"`
	Msg        string `json:"msg"         example:"Ошибка чтения тела запроса"`
}

func NewReadBodyErr() *ReadBodyErr {
	return &ReadBodyErr{CustomCode: 2, Msg: "Ошибка чтения тела запроса"}
}
