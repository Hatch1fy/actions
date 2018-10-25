package actions

import (
	"bytes"
	"time"
)

func parseLine(line []byte) (a Action, msg []byte) {
	separator := bytes.Index(line, separator)
	action := string(line[0:separator])
	a = ParseAction(action)
	msg = line[separator+2:]
	return
}

// Handler is the handler to be called on imports
type Handler func(ts time.Time, a Action, line []byte) error
