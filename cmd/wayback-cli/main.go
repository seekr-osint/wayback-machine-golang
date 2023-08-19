package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/seekr-osint/wayback-machine-golang/internal/wayback"
	"github.com/spf13/cobra"
)

var (
	archiveCmd = &cobra.Command{
		Use:   "archive",
		Short: "Archive the provided URL",
		Run:   archiveRun,
	}
	getSnapshotCmd = &cobra.Command{
		Use:   "get-snapshot",
		Short: "Get the latest existing snapshot",
		Run:   getSnapshotRun,
	}

	rootCmd = &cobra.Command{
		Use:   "wayback",
		Short: "A tool to interact with the Wayback Machine API",
	}
	client = http.Client{}
	output string
)

func main() {
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "pretty", "Output format: json/raw/pretty")
	rootCmd.AddCommand(archiveCmd, getSnapshotCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

func archiveRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Please provide the URL to archive.")
		cmd.Usage()
		return
	}

	url, err := wayback.Archive(args[0], &client)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	switch output {
	case "json":
		fmt.Printf("{\"url\": \"%s\"}\n", url)
	case "raw":
		fmt.Printf("%s\n", url)
	case "pretty":
		fmt.Printf("URL archived at: %s\n", url)
	default:
		fmt.Println("Invalid output option. Please choose 'json', 'raw', or 'pretty'.")
		cmd.Usage()
		return
	}
}

func getSnapshotRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Please provide the snapshot URL.")
		cmd.Usage()
		return
	}

	snapshot, err := wayback.GetSnapshotData(args[0], &client)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	switch output {
	case "json":
		fmt.Printf("%s\n", snapshot)
	case "raw":
		fmt.Printf("%s\n", snapshot.URL)
	case "pretty":
		fmt.Printf("Snapshot found at: %s\n", snapshot.URL)
	default:
		fmt.Println("Invalid output option. Please choose 'json', 'raw', or 'pretty'.")
		cmd.Usage()
		return
	}
}
