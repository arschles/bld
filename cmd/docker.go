package cmd

import (
	"os"
)

func dockerCLIExists() bool {
	if _, err := os.Stat("docker"); err != nil {
		return false
	}
	return true
}
