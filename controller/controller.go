package controller

import "github.com/arizard/pug-go-doro/controller/event"

type Controller interface {
	Quit()
	StartOrPause()
	Skip()
	Listen() event.Event
}
