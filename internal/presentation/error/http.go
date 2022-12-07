package error

// swagger:response HttpError
type HttpError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (he *HttpError) Error() string {
	return he.Message
}
