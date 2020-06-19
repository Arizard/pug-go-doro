package event

type EventKind string

const (
	Quit EventKind = "QUIT"
	StartOrPause   = "START_OR_PAUSE"
	Skip           = "SKIP"
)

type Event struct {
	Kind EventKind
	Numeric float64
	Text string
}
