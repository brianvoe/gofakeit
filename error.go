package gofakeit

type fakeError struct {
	msg string
}

func (f *fakeError) Error() string {
	return f.msg
}

func Error() error {
	return &fakeError{"fake error"}
}
