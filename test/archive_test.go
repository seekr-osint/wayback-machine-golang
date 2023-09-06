package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/seekr-osint/wayback-machine-golang/wayback"
)

func main() {
	client := http.Client{}
	url, err := wayback.Archive("github.com/max", &client) // the clint argument can also be nil
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Archived URL: %s\n", url)
}
