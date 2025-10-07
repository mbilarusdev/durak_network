package network

const (
	// Errors
	CodeServerError   int = 1
	CodeReadBodyError int = 2

	// Client warnings
	CodeUnmarshalingError int = 11
	CodeAlreadyExist      int = 12
	CodeNotFound          int = 13
	CodeWrongInfo         int = 14

	// Token errors
	CodeTokenNotProvided int = 21
	CodeTokenIncorrect   int = 22

	// Success
	CodeSuccess int = 100
)
