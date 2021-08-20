package model

type Exception struct {
	ErrorMessage string
	Error        error
	HTTPCode     int
}

func NewException(errorMessage string, err error, httpCode int) *Exception {
	return &Exception{
		ErrorMessage: errorMessage,
		Error:        err,
		HTTPCode:     httpCode,
	}
}
