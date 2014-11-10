package utils

type SError struct {
	Status int
	Front  error
	Back   error
}

func (e *SError) Error() string {
	return e.Back.Error()
}
