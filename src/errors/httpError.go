package errors

type HttpError struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func UnauthorizedError() HttpError {
	return HttpError{
		401,
		"Unauthorized",
		"You are not authorized to access this resource",
	}

}
