package widget

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

type HistoryWidget struct {
	Name    string
	x, y    int
	w       int
	History []string
}

func (r HistoryWidget) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, _ := g.SetView(r.Name, maxX/2, 0, maxX-1, maxY-5); v != nil {
		v.Clear()
		v.Wrap = true
		v.Title = "History"
		v.Autoscroll = true

		fmt.Fprintf(
			v,
			" \n > %s\n",
			strings.Join(r.History, "\n > "),
		)

	}
	return nil
}

func NewHistoryWidget(name string) *HistoryWidget {
	return &HistoryWidget{Name: name}
}
