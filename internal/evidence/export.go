package evidence

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Report represents a public engineering evidence report.
type Report struct {
	Version   string    `json:"version"`
	Scenario  string    `json:"scenario"`
	Verdict   string    `json:"verdict"`
	Principle string    `json:"principle"`
	SessionID string    `json:"session_id"`
	Transport string    `json:"transport"`
	Timestamp time.Time `json:"timestamp"`
}

// Export writes a public evidence report as JSON.
func Export(filename string, report Report) error {

	report.Timestamp = time.Now().UTC()

	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
