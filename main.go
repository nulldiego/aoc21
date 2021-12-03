package main

import (
	"os"
)

//go:generate go run ./gen
func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
