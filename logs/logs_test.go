package logs

import (
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"kb/machine"
	"os"
	"path"
	"reflect"
	"strings"
	"testing"
	"time"
)

// TestParseLogFolder checks if ParseLogFolder correctly reads a directory of logs and produces the correct report
func TestParseLogFolder(t *testing.T) {
	// given
	// create temporary directory inside the os.TempDir() and a mock log file in it
	tmpFolder, err := ioutil.TempDir(os.TempDir(), "mock_logs")
	tmpFile := createTempLogFile(
		tmpFolder,
		[]byte(`calin tty1                      0:00    0.00s  0:00   0.00s asdf\n`),
	)

	machineName := strings.TrimSuffix(path.Base(tmpFile.Name()), ".log")

	// when
	report, err  := ParseLogFolder(tmpFolder)
	if err != nil {
		t.Error(err)
	}

	// then
	if len(report) != 1 {
		t.Errorf("report should contain 1 log; Contains %d logs", len(report))
	}
	if _, ok := report[machineName]; !ok {
		t.Errorf("LogReport should contain report for: %v", machineName)
	}

	// clean up
	os.Remove(tmpFile.Name())
	os.Remove(tmpFolder)
}

// TestParseLog checks if ParseLogFile correctly parses a given log *os.File of a used computer
func TestParseLog(t *testing.T) {
	// given
	tmpFile := createTempLogFile(
		os.TempDir(),
		[]byte(`calin tty1                      0:00    0.00s  0:00   0.00s asdf\n`),
	)
	defer os.Remove(tmpFile.Name()) // clean up

	// when
	logs := ParseLogFile(tmpFile)

	// then
	if logs.State != machine.Used {
		t.Error("Should report machine full")
	}
}


// TestParseLogEmpty checks if ParseLogFile correctly parses a given log *os.File of a free computer
func TestParseLogEmpty(t *testing.T) {
	// given
	tmpFile := createTempLogFile(
		os.TempDir(),
		[]byte(`calin ssh                      0:00    0.00s  0:00   0.00s asdf\n`),
	)
	defer os.Remove(tmpFile.Name()) // clean up

	// when
	logs := ParseLogFile(tmpFile)

	// then
	if logs.State != machine.Free {
		t.Error("Should report machine empty")
	}
}

// createTempLogFile returns a pointer to a temporary os.File which it creates
func createTempLogFile(dir string, content []byte) *os.File {
	tmpFile, err := ioutil.TempFile(dir, "machine_*.log")
	if err != nil {
		log.Error(err)
	}
	if _, err := tmpFile.Write(content); err != nil {
		log.Error(err)
	}
	if _, err := tmpFile.Seek(io.SeekStart, 0); err != nil {
		log.Error(err)
	}
	return tmpFile
}

func TestDeltaStateChanges(t *testing.T) {
	oldState := machine.LogReport{
		"machine1": {
			State:     machine.Free,
			Timestamp: time.Now(),
		},
	}

	newState := machine.LogReport{
		"machine1": {
			State: machine.Used,
			Timestamp: time.Now(),
		},
	}

	delta := Delta(oldState, newState)
	if !reflect.DeepEqual(delta["machine1"], newState["machine1"]) {
		t.Error("Delta should present new state")
	}
}

func TestDeltaStateDoesNotChange(t *testing.T) {
	mockTime := time.Now()
	oldState := machine.LogReport{
		"machine1": {
			State:     machine.Free,
			Timestamp: mockTime,
		},
	}

	newState := machine.LogReport{
		"machine1": {
			State:     machine.Free,
			Timestamp: mockTime,
		},
	}

	delta := Delta(oldState, newState)
	if len(delta) != 0 {
		t.Errorf("Delta should be empty: %v", delta)
	}
}

func TestStateChangedDifferent(t *testing.T) {
	mockTime := time.Now()

	left := machine.StatusLog{
		State:     machine.Free,
		Timestamp: mockTime.Add(- 1 * time.Hour),
	}

	right := machine.StatusLog{
		State:     machine.Used,
		Timestamp: mockTime,
	}

	if !StateChanged(left, right) {
		t.Error("State should've changed")
	}
}

func TestStateChangedSame(t *testing.T) {
	mockTime := time.Now()

	left := machine.StatusLog{
		State:     machine.Free,
		Timestamp: mockTime,
	}

	right := machine.StatusLog{
		State:     machine.Free,
		Timestamp: mockTime,
	}

	if StateChanged(left, right) {
		t.Error("State should've stayed the same")
	}
}