package machine

import (
	"kb/kilburn"
)

// Status is a representation of the status of a machine after being mapped to a lab
type Status struct {
	StatusLog // get the rest of the info from the log
	Lab       string            `json:"lab"`
}

// Snapshot is a map of machine names to statuses
type Snapshot map[string]Status

// ConvertLogToStatus converts a raw StatusLog into a Status representation
func ConvertLogToStatus(machineName string, log StatusLog) Status {
	return Status{
		log,
		kilburn.MapMachineToLab(machineName),
	}
}

// ConvertReportToSnapshot converts a Report into a Snapshot
func ConvertReportToSnapshot(report LogReport) Snapshot {
	snapshot := make(Snapshot)
	for machine, log := range report {
		snapshot[machine] = ConvertLogToStatus(machine, log)
	}

	return snapshot
}