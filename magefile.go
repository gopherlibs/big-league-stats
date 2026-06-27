//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Test with GoTestSum.
func Test() error {
	return sh.Run("gotestsum", "--", "-coverprofile=coverage.txt", "-covermode=atomic", "./...")
}
