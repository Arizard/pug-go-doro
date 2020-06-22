package core

import (
	"fmt"
	"math"
	"time"

	"github.com/arizard/pug-go-doro/presenter"
	"github.com/arizard/pug-go-doro/status"
)

const (
	defaultDuration     = time.Duration(25 * time.Minute)
	defaultRestDuration = time.Duration(5 * time.Minute)
	tickInterval        = time.Duration(100 * time.Millisecond)
	renderInterval      = time.Duration(100 * time.Millisecond)
)

type Core struct {
	presenter     presenter.Presenter
	ticking       bool
	lastStart     time.Time
	lastDuration  time.Duration
	timeRemaining time.Duration
	status        status.Status
}

func startTick(c *Core) {
	for {
		if c.ticking == true {
			c.tick()
		}
		time.Sleep(tickInterval)
	}
}

func startRender(c *Core) {
	for {
		c.render()
		time.Sleep(renderInterval)
	}
}

func (c *Core) SetDuration(d time.Duration) {
	c.timeRemaining = d
	c.lastDuration = d
	c.presenter.SetTimeRemaining(defaultDuration)
}

func NewCore(pres presenter.Presenter) *Core {

	pres.SetStatus(status.PausedWork)

	c := &Core{
		presenter: pres,
	}

	c.startWorkPeriod()
	c.StartOrPause()

	go startTick(c)
	go startRender(c)

	return c
}

func (c *Core) setStatus(s status.Status) {
	c.status = s
	c.presenter.SetStatus(s)
}

func (c *Core) startWorkPeriod() {
	c.setStatus(status.RunningWork)
	c.SetTicking(true)
	c.SetDuration(defaultDuration)
	c.lastStart = time.Now()
}

func (c *Core) startRestPeriod() {
	c.setStatus(status.RunningRest)
	c.SetTicking(true)
	c.SetDuration(defaultRestDuration)
	c.lastStart = time.Now()
}

func (c *Core) StartOrPause() {
	c.SetTicking(!c.ticking)

	if c.status == status.RunningWork {
		c.setStatus(status.PausedWork)
	} else if c.status == status.RunningRest {
		c.setStatus(status.PausedRest)
	} else if c.status == status.PausedWork {
		c.setStatus(status.RunningWork)
	} else if c.status == status.PausedRest {
		c.setStatus(status.RunningRest)
	}

}

func (c *Core) SetTicking(shouldTick bool) {
	c.ticking = shouldTick
}

func (c *Core) Skip() {
	c.presenter.AddHistory(
		fmt.Sprintf("Finished %s work, (%s)", (c.lastDuration - c.timeRemaining).String(), c.lastStart.Format("3:04 PM")),
	)
	time.Sleep(50 * time.Millisecond)
	c.startWorkPeriod()
}

func (c *Core) consumeTime(d time.Duration) {
	c.timeRemaining = c.timeRemaining - d
	c.timeRemaining = time.Duration(math.Max(0, float64(c.timeRemaining)))
}

func (c *Core) tick() {
	c.consumeTime(tickInterval)

	if c.timeRemaining == 0 && c.ticking == true {
		if c.status == status.RunningWork {
			c.StartOrPause()
			c.presenter.AddHistory(
				fmt.Sprintf("Finished %s work, (%s)", c.lastDuration.String(), c.lastStart.Format("3:04 PM")),
			)
			c.startRestPeriod()
		} else if c.status == status.RunningRest {
			c.StartOrPause()
			c.presenter.AddHistory(
				fmt.Sprintf("Finished %s rest, (%s)", c.lastDuration.String(), c.lastStart.Format("3:04 PM")),
			)

			c.startWorkPeriod()
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func (c *Core) render() {
	c.presenter.SetTimeRemaining(c.timeRemaining)
}
