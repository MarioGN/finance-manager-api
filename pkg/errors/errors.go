package errors

type ApplicationError struct {
	ErrorMessage string `json:"error_message"`
}

func NewApplicationError(message string) ApplicationError {
	return ApplicationError{ErrorMessage: message}
}

var InternnalServerError = NewApplicationError("internal server error")
var NotFoundError = NewApplicationError("not found")
var InvalidRequestError = NewApplicationError("invalid request payload")
