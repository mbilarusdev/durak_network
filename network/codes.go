package network

const (
	// Errors
	CodeServerError       int = 1
	CodeUnmarshalingError int = 2

	// Client warnings
	CodeIncorrectBody int = 11
	CodeAlreadyExist  int = 12
	CodeNotFound      int = 13

	// Token errors
	CodeTokenNotProvided int = 21
	CodeTokenIncorrect   int = 22

	// Success
	CodeSuccess int = 100
)
