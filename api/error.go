package api

func (e *Error) Error() string {
	return *(e.Message)
}
