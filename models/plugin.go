package models

type Plugin interface {
	Info() string
	Transform(logs *ForscanLogs) error
}
