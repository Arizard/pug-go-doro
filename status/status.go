package status

type Status string

const (
	Stopped     Status = "STOPPED"
	RunningWork        = "RUNNING_WORK"
	RunningRest        = "RUNNING_REST"
	PausedWork         = "PAUSED_WORK"
	PausedRest         = "PAUSED_REST"
)

var statusNameByStatus = map[Status]string{
	Stopped: "Stopped",
	RunningWork: "Working",
	RunningRest: "Resting",
	PausedWork: "Paused (Working)",
	PausedRest: "Paused (Resting)",
}

func (s Status) String() string {
	if name, ok := statusNameByStatus[s]; ok {
		return name
	}
	return string(s)
}
