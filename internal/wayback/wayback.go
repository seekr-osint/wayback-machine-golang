package wayback

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrArchivedUrl = errors.New("can't get archived url")
)

// Archive creates an archive of the given URL using the Wayback Machine.
// It takes a URL and an optional http.Client as parameters.
// If the client is nil, a default client will be used.
// It returns the URL of the archived snapshot and any error encountered.
func Archive(url string, client *http.Client) (string, error) {
	if client == nil {
		client = &http.Client{}
	}

	response, err := client.Get(fmt.Sprintf("https://web.archive.org/save/%s", url))
	if err != nil {
		return "", err
	}
	return response.Request.URL.String(), nil
}

// GetSnapshotData retrieves snapshot data for the given URL from the Wayback Machine.
// It takes a URL and an http.Client as parameters.
// It returns a pointer to SnapshotData and any error encountered.
func GetSnapshotData(url string, client *http.Client) (*SnapshotData, error) {
	response, err := client.Get(fmt.Sprintf("https://archive.org/wayback/available?url=%s", url))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var snapshotData SnapshotData
	err = json.NewDecoder(response.Body).Decode(&snapshotData)
	if err != nil {
		return nil, err
	}

	return &snapshotData, nil
}
