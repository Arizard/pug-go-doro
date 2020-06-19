package widget

import (
	"fmt"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/jroimartin/gocui"

	"github.com/arizard/pug-go-doro/status"
)

type ControlsWidget struct {
	Name          string
	x, y          int
	w             int
	Status        status.Status
	TimeRemaining time.Duration
}

func durationAs4DigitWithColon (dur time.Duration) string {
	totalSeconds := dur.Seconds()

	var minutes = int(totalSeconds / 60)
	var seconds = int(totalSeconds) % 60

	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

func (c ControlsWidget) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if mainView, _ := g.SetView(c.Name, 0, 0, maxX/2 - 1, maxY - 5); mainView != nil {
		mainView.Clear()
		mainView.Title = "Pug Go Doro!"

		timeRemainingContent := durationAs4DigitWithColon(c.TimeRemaining)

		heading := figure.NewFigure(
			fmt.Sprintf("%s", timeRemainingContent),
			"",
			true,
		)

		fmt.Fprintf(
			mainView,
			" \n%s\n > Time Remaining: %s\n > Status: %s\n",
			heading.String(),
			timeRemainingContent,
			c.Status,
		)
	}

	helpTextContent := "    [SPACE] Start/Pause    [CTRL+SPACE] Skip    [CTRL+C] Quit    "
	helpTextContentLength := len(helpTextContent) + len(helpTextContent) % 2

	if helpText, _ := g.SetView("helptext", (maxX/2 - 1) - helpTextContentLength/2, maxY - 4, (maxX/2 - 1) + helpTextContentLength/2, maxY - 2); helpText != nil {
		helpText.Clear()
		fmt.Fprintf(
			helpText,
			helpTextContent,
		)
	}
	return nil
}

func NewControlsWidget(name string) *ControlsWidget {
	return &ControlsWidget{Name: name}
}
