package cmd

import (
	"io/ioutil"

	"github.com/magefile/mage/sh"
)

func azCLIExists() bool {
	if _, err := sh.Exec(map[string]string{}, ioutil.Discard, ioutil.Discard, "az"); err != nil {
		return false
	}
	return true
}
