package dto

type ErrorDto struct {
	Error string `json:"error"`
}

func Error(msg string) *ErrorDto {
	return &ErrorDto{
		Error: msg,
	}
}