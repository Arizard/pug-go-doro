package controller

import (
	"github.com/arizard/pug-go-doro/controller/event"
)

type Default struct {
	c chan event.Event
}

func (d Default) StartOrPause() {
	d.c <- event.Event{
		Kind: event.StartOrPause,
		Numeric: 0,
		Text: "start or pause from controller",
	}
}

func (d Default) Skip() {
	d.c <- event.Event{
		Kind: event.Skip,
		Numeric: 0,
		Text: "start or pause from controller",
	}
}

func (d Default) Quit() {
	d.c <- event.Event{
		Kind: event.Quit,
		Numeric: 0,
		Text: "quit from controller",
	}
}

func (d Default) Listen() event.Event {
	return <- d.c
}

func NewDefault() *Default {
	return &Default{
		c: make(chan event.Event),
	}
}
