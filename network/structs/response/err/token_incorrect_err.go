package net_err

type TokenIncorrectErr struct {
	CustomCode int    `json:"custom_code" example:"22"`
	Msg        string `json:"msg"         example:"Токен инкорректен"`
}

func NewTokenIncorrectErr() *TokenIncorrectErr {
	return &TokenIncorrectErr{CustomCode: 22, Msg: "Токен инкорректен"}
}
