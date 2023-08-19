package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/seekr-osint/wayback-machine-golang/internal/wayback"
)

func main() {
	client := http.Client{}
	url, err := wayback.Archive("github.com/max", &client)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("archived url: %s\n", url)
}
