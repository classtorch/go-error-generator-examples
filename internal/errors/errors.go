package errors

func (e *Error) Error() string {
	return e.Msg
}