package templates

import "github.com/gdbu/poller"

func isRelevantEvent(evt poller.Event) (ok bool) {
	switch evt {
	case poller.EventWrite:
	case poller.EventChmod:
	case poller.EventCreate:

	default:
		return false
	}

	return true
}
