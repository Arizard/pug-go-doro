package presenter

import (
	"time"

	"github.com/arizard/pug-go-doro/status"
)

type Presenter interface {
	SetTimeRemaining(dur time.Duration)
	SetStatus(s status.Status)
	AddHistory(s string)
	Close()
}