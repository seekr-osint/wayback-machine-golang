package wayback

import (
	"time"
)

const customTimeLayout = "20060102150405"

type ArchivedSnapshot struct {
	Status    string    `json:"status"`
	Available bool      `json:"available"`
	URL       string    `json:"url"`
	Timestamp time.Time `json:"timestamp"`
}

type SnapshotData struct {
	URL               string `json:"url"`
	ArchivedSnapshots struct {
		Closest ArchivedSnapshot `json:"closest"`
	} `json:"archived_snapshots"`
}
