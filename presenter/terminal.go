package presenter

import (
	"log"
	"time"

	"github.com/jroimartin/gocui"

	"github.com/arizard/pug-go-doro/controller"
	"github.com/arizard/pug-go-doro/notify"
	"github.com/arizard/pug-go-doro/presenter/widget"
	"github.com/arizard/pug-go-doro/status"
)

type Terminal struct {
	g              *gocui.Gui
	ctrl           controller.Controller
	HistoryWidget  *widget.HistoryWidget
	ControlsWidget *widget.ControlsWidget
}

func (t *Terminal) AddHistory(rec string) {
	t.g.Update(func(g *gocui.Gui) error {
		t.HistoryWidget.History = append(t.HistoryWidget.History, rec)
		return nil
	})

	notify.Info(rec)
}

func (t *Terminal) SetTimeRemaining(dur time.Duration) {
	t.g.Update(func(g *gocui.Gui) error {
		t.ControlsWidget.TimeRemaining = dur
		return nil
	})
}

func (t *Terminal) SetStatus(s status.Status) {
	t.g.Update(func(g *gocui.Gui) error {
		t.ControlsWidget.Status = s
		return nil
	})
}

func mainLoop(g *gocui.Gui) {
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func setKeyBindings(g *gocui.Gui, t *Terminal) {
	g.SetKeybinding(
		"",
		gocui.KeyCtrlC,
		gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			t.ctrl.Quit()
			return nil
		},
	)

	g.SetKeybinding(
		"",
		gocui.KeyCtrlL,
		gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			t.AddHistory("00m 00s Test at <Now>")
			return nil
		},
	)

	g.SetKeybinding(
		"",
		gocui.KeySpace,
		gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			t.ctrl.StartOrPause()
			return nil
		},
	)

	g.SetKeybinding(
		"",
		gocui.KeyCtrlSpace,
		gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			t.ctrl.Skip()
			return nil
		},
	)
}

func (t *Terminal) Close() {
	t.g.Close()
}

func NewTerminal(ctrl controller.Controller) *Terminal {
	g, err := gocui.NewGui(gocui.OutputNormal)

	t := &Terminal{g: g}

	if err != nil {
		log.Panicln(err)
	}

	t.HistoryWidget = widget.NewHistoryWidget("History")
	t.ControlsWidget = widget.NewControlsWidget("Controls")

	g.SetManager(t.HistoryWidget, t.ControlsWidget)

	setKeyBindings(g, t)

	go mainLoop(g)

	t.ctrl = ctrl

	return t
}
