package simply

import (
	"encoding/json"
	"log"
)

type F = map[string]interface{}

type Logger interface {
	With(fields F) Logger
	Err(err error) Logger

	Info(msg string)
	Error(msg string)
}

func NewLogger() Logger {
	return logger{
		F{},
	}
}

type logger struct {
	kvs F
}

func (l logger) With(fields F) Logger {
	for k, v := range fields {
		l.kvs[k] = v
	}

	return l
}

func (l logger) Info(msg string) {
	l.kvs["msg"] = msg
	l.kvs["level"] = "info"

	jsonStr, err := json.Marshal(l.kvs)
	if err != nil {
		log.Print(l.kvs)
	} else {
		log.Print(string(jsonStr))
	}
}

func (l logger) Error(msg string) {
	l.kvs["msg"] = msg
	l.kvs["level"] = "error"

	jsonStr, err := json.Marshal(l.kvs)
	if err != nil {
		log.Print(l.kvs)
	} else {
		log.Print(string(jsonStr))
	}
}

func (l logger) Err(err error) Logger {
	l.kvs["err"] = err

	return l
}
