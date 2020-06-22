package logger

type Logger interface {
	Write(a ...interface{})
	Error(description string, err error)
}
