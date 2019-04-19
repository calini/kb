package machine

import (
	"testing"
	"time"
)

// TestGetStatus checks that conversion is done right
func TestGetStatus(t *testing.T) {
	statusLog := StatusLog{
		Used,
		time.Now(),
	}

	status := ConvertLogToStatus("retina", statusLog)
	if status.Lab != "none" {
		t.Error("Should be reported as none")
	}

	if statusLog.State != status.State {
		t.Error("Should preserve rest of object")
	}
}
