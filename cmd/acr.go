package cmd

import (
	"os"
)

func azCLIExists() bool {
	if _, err := os.Stat("az"); err != nil {
		return false
	}
	return true
}
