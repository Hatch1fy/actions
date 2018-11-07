package actions

import (
	"bytes"
	"time"
)

func parseLine(line []byte) (a Action, key, value []byte) {
	// Find first separator
	sepIndex := bytes.Index(line, separator)
	// Parse action
	action := string(line[0:sepIndex])
	a = ParseAction(action)

	// Remove action and separator
	line = line[sepIndex+2:]

	// Find second separator
	sepIndex = bytes.Index(line, separator)
	// Set key as bytes until separator
	key = line[0:sepIndex]

	// Remove key and separator
	line = line[sepIndex+2:]

	// Set value as remaining bytes
	value = line
	return
}

// Handler is the handler to be called on imports
type Handler func(ts time.Time, a Action, key, value []byte) error
