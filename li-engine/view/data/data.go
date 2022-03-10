package data

import "fmt"

type Handler func(variables interface{}) (res interface{}, err error)

var handlers = make(map[string]Handler)

func RegisterHandler(name string, handler Handler) {
	handlers[name] = handler
}

func ExecHandler(name string, variables interface{}) (res interface{}, err error) {
	handler, exists := handlers[name]
	if !exists {
		return nil, fmt.Errorf("handler `%s` not register", name)
	}
	return handler(variables)
}
