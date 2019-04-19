package logs

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"kb/machine"
)

// CronFrequency is the expected frequency of the cron that produces the log files
const CronFrequency = 1 * time.Minute

// ParseLogFolder returns the state representation based on the logs folder provided
func ParseLogFolder(logsPath string) (machine.LogReport, error) {
	// Read all files in directory
	fileDescriptions, err := ioutil.ReadDir(logsPath)
	if err != nil {
		return nil, err
	}

	// Get the machines status
	machines := make(machine.LogReport)
	for _, fileInfo := range fileDescriptions {
		file, err := os.Open(path.Join(logsPath, fileInfo.Name()))
		if err != nil {
			fmt.Errorf("Unable to open file: %v", err)
		}
		defer file.Close()

		machineName := strings.TrimSuffix(fileInfo.Name(), ".log")
		machines[machineName] = ParseLogFile(file)
	}

	return machines, nil
}

// ParseLogFile parses a log file produced by "w -h" and returns a machine.StatusLog
func ParseLogFile(file *os.File) machine.StatusLog {
	scanner := bufio.NewScanner(file)

	userPresent := false
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Fields(text)

		// filter out empty lines
		if len(fields) < 2 {
			continue
		}

		// if we find any tty line
		if strings.Contains(fields[1], "tty") {
			userPresent = true
		}
	}

	// get the file info
	stats, err := file.Stat()
	if err != nil {
		fmt.Errorf("Unable to read file information: %v", err)
	}

	// Determine state based on file modTime and content
	var state machine.State
	if time.Now().Add(-1 * CronFrequency).After(stats.ModTime()) { // if log is older than reference time, machine did not report so is offline
		state = machine.Offline
	} else if userPresent {
		state = machine.Used
	} else {
		state = machine.Free
	}

	// Return machine status
	return machine.StatusLog{
		State:     state,
		Timestamp: stats.ModTime(),
	}
}

// Delta gives the status in difference (the minimum that has to be sent for the state to be reflected at the other side)
func Delta(left, right machine.LogReport) machine.LogReport {
	diff := make(machine.LogReport)
	// for each of the new ones
	for name := range right {
		if StateChanged(left[name], right[name]) {
			diff[name] = right[name]
		}
	}

	return diff
}

// StateChanged checks for state changes between two StatusLogs
func StateChanged(left, right machine.StatusLog) bool {
	return left.State != right.State
}
