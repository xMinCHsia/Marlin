
// Package config loads and validates the Marlin configuration document.
package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Target holds the target-engine binding.
type Target struct {
