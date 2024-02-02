package xor_error

type IllegalArgumentError struct {
	message string
}

func NewIllegalArgumentError(message string) error {
	return IllegalArgumentError{message: message}
}

func (r IllegalArgumentError) Error() string {
	return r.message
}
