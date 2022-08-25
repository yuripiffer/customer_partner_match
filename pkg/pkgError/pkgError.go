package pkgError

type ErrorStruct struct {
	AppError
	Err      string
	Message  string
	ErrorKey string
}

type AppError interface {
	error
	GetErr() string
	GetMessage() string
	GetErrorKey() string
}

func (e *ErrorStruct) GetErr() string {
	return e.Err
}

func (e *ErrorStruct) GetMessage() string {
	return e.Message
}

func (e *ErrorStruct) GetErrorKey() string {
	return e.ErrorKey
}

const (
	ServerError = "error.InternalServerError"
	InputError  = "error.InputError"
)

func NewServerError(message string, err error) AppError {
	return &ErrorStruct{
		Err:      err.Error(),
		Message:  message,
		ErrorKey: ServerError,
	}
}

func NewInputError(message string, err error) AppError {
	return &ErrorStruct{
		Err:      err.Error(),
		Message:  message,
		ErrorKey: InputError,
	}
}
