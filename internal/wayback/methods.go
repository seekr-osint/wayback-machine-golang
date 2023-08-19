package wayback

import (
	"encoding/json"
	"fmt"
	"time"
)

func (a ArchivedSnapshot) MarshalJSON() ([]byte, error) {
	type Alias ArchivedSnapshot
	return json.Marshal(&struct {
		Alias
		ArchivedSnapshot string `json:"timestamp"`
	}{
		Alias:            (Alias)(a),
		ArchivedSnapshot: a.Timestamp.Format(customTimeLayout),
	})
}

func (a *ArchivedSnapshot) UnmarshalJSON(data []byte) error {
	type Alias ArchivedSnapshot
	aux := &struct {
		*Alias
		ArchivedSnapshot string `json:"timestamp"`
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	parsedTime, err := time.Parse(customTimeLayout, aux.ArchivedSnapshot)
	if err != nil {
		return err
	}
	a.Timestamp = parsedTime
	return nil
}

func (a ArchivedSnapshot) String() string {
	jsonString, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error marshaling JSON: %v", err)
	}
	return string(jsonString)
}

func (s SnapshotData) String() string {
	jsonString, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error marshaling JSON: %v", err)
	}
	return string(jsonString)
}
