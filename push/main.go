package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"kb/logs"
	"net/http"
	"os"
	"time"

	"kb/machine"

	log "github.com/sirupsen/logrus"
)

const (
	frequency = 1 * time.Minute // Is in sync with cron frequency
	pushURL   = "http://localhost:8080/report"
)

func init() {
	// Add timestamp to logs
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)
}

func main() {
	// Get logs path
	if len(os.Args) != 2 {
		log.Fatal("You must provide the logs path!")
	}
	logsPath := os.Args[1]
	if _, err := os.Stat(logsPath); os.IsNotExist(err) {
		log.Fatal("Path does not exist!")
	}

	// Send initial state of the system
	snapshot, _ := logs.ParseLogFolder(logsPath)
	for {
		err := sendReport(snapshot)
		if err != nil {
			log.Errorf("could not send initial report: %s; retrying...", err.Error())
			time.Sleep(frequency)
		} else {
			log.Infof("initial report sent")
			break
		}
	}

	// After that, only send deltas
	ticker := time.NewTicker(frequency)
	for {
		select {
		case <-ticker.C:
			newSnapshot, _ := logs.ParseLogFolder(logsPath)
			delta := logs.Delta(snapshot, newSnapshot)

			if err := sendReport(delta); err != nil {
				log.Errorf("could not send delta report: %s; retrying...", err.Error())
			} else {
				log.Infof("delta report sent")
			}
			snapshot = newSnapshot
		}
	}
}

// sendReport sends a report to the API
func sendReport(report machine.LogReport) error {

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(report)

	resp, err := http.Post(pushURL, "application/json", body)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		return errors.New("server did not respond with status code 200")
	}
	return nil
}
