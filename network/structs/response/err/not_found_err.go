package net_err

type NotFoundErr struct {
	CustomCode int    `json:"custom_code" example:"13"`
	Msg        string `json:"msg"         example:"Сущность не найдена"`
}

func NewNotFoundErr() *NotFoundErr {
	return &NotFoundErr{CustomCode: 13, Msg: "Сущность не найдена"}
}
