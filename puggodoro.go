package main

import (
	"time"

	"github.com/arizard/pug-go-doro/controller"
	event "github.com/arizard/pug-go-doro/controller/event"
	"github.com/arizard/pug-go-doro/core"
	"github.com/arizard/pug-go-doro/presenter"
)

func main() {
	var ctrl controller.Controller = controller.NewDefault()

	var pres presenter.Presenter = presenter.NewTerminal(ctrl)
	defer pres.Close()

	var core *core.Core = core.NewCore(pres)

	for {
		ctrlEvent := ctrl.Listen()

		if ctrlEvent.Kind == event.Quit {
			break
		}

		if ctrlEvent.Kind == event.StartOrPause {
			core.StartOrPause()
		}

		if ctrlEvent.Kind == event.Skip {
			core.Skip()
		}

		time.Sleep(time.Duration(5 * time.Millisecond))
	}
}
