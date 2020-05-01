package lib

type splitError struct {
	msg string
}

func (se splitError) Error() string {
	return "foo"
}

func NewSplitError(s string) splitError {
	var se splitError
	se.msg = s
	return se
}

type initError struct {
}

func (ie *initError) Error() string {
	return "foo"
}
