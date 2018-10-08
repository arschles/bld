package cmd

import (
	"io/ioutil"

	"github.com/magefile/mage/sh"
)

func dockerCLIExists() bool {
	if _, err := sh.Exec(map[string]string{}, ioutil.Discard, ioutil.Discard, "docker"); err != nil {
		return false
	}
	return true
}
