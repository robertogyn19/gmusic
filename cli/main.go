package main

import (
	"log"

	"github.com/robertogyn19/gmusic/cli/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalf("could not execute command, error: %v", err)
	}
}
