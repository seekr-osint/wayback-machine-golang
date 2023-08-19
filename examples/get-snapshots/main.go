package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/seekr-osint/wayback-machine-golang/wayback"
)

func main() {
	client := http.Client{}
	snapshots, err := wayback.GetSnapshotData("github.com/max", &client)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("snapshot data: %s\n", snapshots)
}
