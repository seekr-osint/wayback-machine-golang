package wayback

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/seekr-osint/wayback-machine-golang/internal/wayback"
)

func TestArchivedSnapshotJSONMarshal(t *testing.T) {
	snapshot := wayback.ArchivedSnapshot{
		Status:    "200",
		Available: true,
		URL:       "http://web.archive.org/web/20230819204825/https://github.com/greg",
		Timestamp: time.Date(2023, 8, 19, 20, 48, 25, 0, time.UTC),
	}

	expectedJSON := `{"status":"200","available":true,"url":"http://web.archive.org/web/20230819204825/https://github.com/greg","timestamp":"20230819204825"}`

	resultJSON, err := json.Marshal(snapshot)
	if err != nil {
		t.Errorf("Error marshaling JSON: %v", err)
	}

	if string(resultJSON) != expectedJSON {
		t.Errorf("Expected JSON:\n%s\nActual JSON:\n%s", expectedJSON, resultJSON)
	}
}

func TestArchivedSnapshotJSONUnmarshal(t *testing.T) {

	jsonData := []byte(`{"status":"200","available":true,"url":"http://web.archive.org/web/20230819204825/https://github.com/greg","timestamp":"20230819204825"}`)

	var snapshot wayback.ArchivedSnapshot

	err := json.Unmarshal(jsonData, &snapshot)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	expectedTimestamp := time.Date(2023, 8, 19, 20, 48, 25, 0, time.UTC)
	if !snapshot.Timestamp.Equal(expectedTimestamp) {
		t.Errorf("Expected timestamp: %v\nActual timestamp: %v", expectedTimestamp, snapshot.Timestamp)
	}
}

func TestArchivedSnapshotRoundTrip(t *testing.T) {
	snapshot := wayback.ArchivedSnapshot{
		Status:    "200",
		Available: true,
		URL:       "http://web.archive.org/web/20230819204825/https://github.com/greg",
		Timestamp: time.Date(2023, 8, 19, 20, 48, 25, 0, time.UTC),
	}

	resultJSON, err := json.Marshal(snapshot)
	if err != nil {
		t.Errorf("Error marshaling JSON: %v", err)
	}

	var unmarshaledSnapshot wayback.ArchivedSnapshot

	err = json.Unmarshal(resultJSON, &unmarshaledSnapshot)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	if snapshot.Status != unmarshaledSnapshot.Status ||
		snapshot.Available != unmarshaledSnapshot.Available ||
		snapshot.URL != unmarshaledSnapshot.URL ||
		!snapshot.Timestamp.Equal(unmarshaledSnapshot.Timestamp) {
		t.Errorf("Snapshot instances are not equal:\nOriginal: %+v\nUnmarshaled: %+v", snapshot, unmarshaledSnapshot)
	}
}
