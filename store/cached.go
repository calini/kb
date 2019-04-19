package store

import "kb/machine"

// CacheStore stores the machine Snapshot in memory and has query methods for it
type CacheStore struct {
	cache machine.Snapshot // you might need to move this and use make()
}

// NewCacheStore initializes and returns a new CacheStore
func NewCacheStore() CacheStore {
	return CacheStore{make(machine.Snapshot)}
}

// SaveSnapshot updates the state of the snapshot
func (s *CacheStore) SaveSnapshot(report machine.Snapshot) {
	// Update the latest state per machine
	for machine := range report {
		s.cache[machine] = report[machine]
	}
}

// GetSnapshot returns the status set of all machines
func (s *CacheStore) GetSnapshot() machine.Snapshot {
	return s.cache
}

// GetSnapshotForMachine returns the status of one machine
func (s *CacheStore) GetSnapshotForMachine(machineName string) machine.Status {
	return s.cache[machineName]
}

// GetSnapshotForLab returns the status set of all machines in a lab
func (s *CacheStore) GetSnapshotForLab(labName string) machine.Snapshot {
	machines := make(machine.Snapshot)

	for machine, stat := range s.cache {
		if stat.Lab == labName {
			machines[machine] = stat
		}
	}

	return machines
}
