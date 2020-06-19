package notify

import "github.com/martinlindhe/notify"

const defaultIcon = ""

func Info(message string) {
	notify.Notify("Puggodoro", "Info", message, defaultIcon)
}

func Alert(message string) {
	notify.Notify("Puggodoro", "Alert", message, defaultIcon)
}

