package utils

type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErr(code int, message string) *Err {
	return &Err{
		Code:    code,
		Message: message,
	}
}
func (e *Err) Error() string {
	return e.Message
}
func NewErrWithMessage(message string) *Err {
	return &Err{
		Code:    500,
		Message: message,
	}
}
