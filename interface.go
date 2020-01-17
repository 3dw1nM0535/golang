package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

// Widget is an interface
type Widget interface {
	// ID returns a widget's unique identifier
	ID() string
}

type widget struct {
	id string
}

// NewWidget() return a new widget
func NewWidget() widget {
	return widget{
		id: uuid.NewV4().String(),
	}
}

func (w widget) ID() string {
	return w.id
}

func main() {
	w := Widget.ID(NewWidget())
	fmt.Println(w)
}
