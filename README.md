# Wayback Machine Golang Library

The Wayback Machine Golang Library is a Go package that provides a simple and convenient way to interact with the Wayback Machine's API to archive and retrieve snapshot data of web pages. The library is designed to streamline the process of archiving URLs and fetching historical snapshots from the Wayback Machine.

## Usage

### Archiving a URL

This example demonstrates how to use the library to archive a URL using the Wayback Machine API.

```go
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
```

### Retrieving Snapshot Data

This example showcases how to retrieve snapshot data of a URL from the Wayback Machine API.

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/seekr-osint/wayback-machine-golang/wayback"
)

func main() {
	client := http.Client{}
	snapshots, err := wayback.GetSnapshotData("github.com/max", &client) // the clint argument can also be nil
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Snapshot data: %s\n", snapshots)
}
```

## License

This library is distributed under the [GPL v3 License](LICENSE).
