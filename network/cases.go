package network

import (
	"net/http"
)

func Success[T any](w http.ResponseWriter, res T) *DurakHandlerResult {
	return (&DurakResponse[T]{
		Code: http.StatusOK,
		Data: res,
	}).Send(w)
}

func SuccessString(w http.ResponseWriter, res string) *DurakHandlerResult {
	return (&DurakResponse[string]{
		Code: http.StatusOK,
		Data: res,
	}).Send(w)
}

func SuccessJWT(w http.ResponseWriter, jwt string) *DurakHandlerResult {
	return (&DurakResponse[string]{
		Code: http.StatusOK,
		Jwt:  jwt,
	}).Send(w)
}

func AlreadyExistJWT(w http.ResponseWriter, jwt string) *DurakHandlerResult {
	return (&DurakResponse[string]{
		Code: http.StatusOK,
		Jwt:  jwt,
		Err:  &ErrorPart{Code: CodeAlreadyExist, Msg: "У вас уже есть токен"},
	}).Send(w)
}

func ReadBodyError(w http.ResponseWriter) *DurakHandlerResult {
	return (&DurakResponse[any]{
		Code: http.StatusBadRequest,
		Err:  &ErrorPart{Code: CodeReadBodyError, Msg: "Ошибка чтения Body"},
	}).Send(w)
}

func UnmarshalingError(w http.ResponseWriter) *DurakHandlerResult {
	return (&DurakResponse[any]{
		Code: http.StatusBadRequest,
		Err:  &ErrorPart{Code: CodeUnmarshalingError, Msg: "Ошибка анмаршалинга Body"},
	}).Send(w)
}

func NotFound(w http.ResponseWriter, msg string) *DurakHandlerResult {
	return (&DurakResponse[any]{
		Code: http.StatusNotFound,
		Err:  &ErrorPart{Code: CodeNotFound, Msg: msg},
	}).Send(w)
}

func TokenNotProvided(w http.ResponseWriter) *DurakHandlerResult {
	return (&DurakResponse[any]{
		Code: http.StatusUnauthorized,
		Err:  &ErrorPart{Code: CodeTokenNotProvided, Msg: "Токен не передан"},
	}).Send(w)
}

func TokenIncorrect(w http.ResponseWriter) *DurakHandlerResult {
	return (&DurakResponse[any]{
		Code: http.StatusUnauthorized,
		Err:  &ErrorPart{Code: CodeTokenIncorrect, Msg: "Токен неверен"},
	}).Send(w)
}

func WrongInfo(w http.ResponseWriter, msg string) *DurakHandlerResult {
	return (&DurakResponse[any]{
		Code: http.StatusBadRequest,
		Err:  &ErrorPart{Code: CodeWrongInfo, Msg: msg},
	}).Send(w)
}

func ServerError(w http.ResponseWriter, msg string) *DurakHandlerResult {
	return (&DurakResponse[any]{
		Code: http.StatusInternalServerError,
		Err:  &ErrorPart{Code: CodeServerError, Msg: msg},
	}).Send(w)
}
