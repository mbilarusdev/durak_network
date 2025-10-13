package net_err

type TokenNotProvidedErr struct {
	CustomCode int    `json:"custom_code" example:"21"`
	Msg        string `json:"msg"         example:"Токен не передан"`
}

func NewTokenNotProvidedErr() *TokenNotProvidedErr {
	return &TokenNotProvidedErr{CustomCode: 21, Msg: "Токен не передан"}
}
