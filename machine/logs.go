package machine

import "time"

// State is a representation of the state of the machine (free, used, offline)
type State string
const (
	// Free means machine is up for grabs
	Free State = "free"
	// Used means someone is at the machine
	Used State = "used"
	// Offline means machine is turned off or not responding (maybe it's in Windows mode)
	Offline State = "offline"
)

// StatusLog is a representation of the status of a machine as received from the producer
type StatusLog struct {
	State     State             `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
}

// LogReport is a map of machine names to status logs
type LogReport map[string]StatusLog
