package net_res

type Body struct {
	Code int `json:"code"`
	Err  any `json:"err,omitempty"`
	Data any `json:"data,omitempty"`
}

func NewBody(code int, err any, data any) *Body {
	body := new(Body)
	body.Code = code
	body.Err = err
	body.Data = data

	return body
}
