package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "boxstash",
    Short: "Server and client for privately hosting vagrant boxes",
    Long:  `Boxstash provides an API server for privately hosting custom vagrant boxes on a Vagrant Cloud compatible service, and a CLI client`,
    Run: func(cmd *cobra.Command, args []string) {
        runClient()
    },
}

func Execute() error {
    return rootCmd.Execute()
}

func runClient() {
    fmt.Println("See `boxstash -h` for more info.  " +
        "This utility will be a full client of the server api.")
}