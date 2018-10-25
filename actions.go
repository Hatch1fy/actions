package actions

import (
	"encoding/json"

	"github.com/Hatch1fy/logger"
)

var (
	separator = []byte("::")
)

const (
	// Break will end a ForEach loop early without returning an error
	// Note: This is an alias of logger.Break
	Break = logger.Break
)

// New will return a new instance of Actions
func New(dir, name string) (ap *Actions, err error) {
	var a Actions
	if a.Logger, err = logger.New(dir, name); err != nil {
		return
	}

	ap = &a
	return
}

// Actions manages actions
type Actions struct {
	*logger.Logger
}

// Log will log an action with a byteslice message
func (a *Actions) Log(action Action, msg []byte) (err error) {
	// Create byteslice with action string
	bs := []byte(action.String())
	// Append separator
	bs = append(bs, separator...)
	// Append message
	bs = append(bs, msg...)
	return a.Logger.Log(bs)
}

// LogString will log an action with a string message
func (a *Actions) LogString(action Action, msg string) (err error) {
	return a.Log(action, []byte(msg))
}

// LogJSON will log an action with a JSON message
func (a *Actions) LogJSON(action Action, msg interface{}) (err error) {
	var bs []byte
	if bs, err = json.Marshal(msg); err != nil {
		return
	}

	return a.Log(action, bs)
}

// Close will close an instance of Actions
func (a *Actions) Close() (err error) {
	return a.Logger.Close()
}
