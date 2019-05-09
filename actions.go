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

func (a *Actions) logEntries(entries []transactionEntry) (err error) {
	for _, entry := range entries {
		if err = a.Log(entry.action, entry.key, entry.value); err != nil {
			return
		}
	}

	return
}

// Log will log an action with a byteslice message
func (a *Actions) Log(action Action, key, value []byte) (err error) {
	// Create byteslice with action string
	bs := []byte(action.String())
	// Append separator
	bs = append(bs, separator...)
	// Append key
	bs = append(bs, key...)
	// Append separator
	bs = append(bs, separator...)
	// Append value
	bs = append(bs, value...)
	return a.Logger.Log(bs)
}

// LogString will log an action with a string message
func (a *Actions) LogString(action Action, key, value string) (err error) {
	return a.Log(action, []byte(key), []byte(value))
}

// LogJSON will log an action with a JSON message
func (a *Actions) LogJSON(action Action, key []byte, value interface{}) (err error) {
	var bs []byte
	if bs, err = json.Marshal(value); err != nil {
		return
	}

	return a.Log(action, key, bs)
}

// Transaction will initialize a new logging transaction
// Note: Logs will not flush until the transaction is closed without error
func (a *Actions) Transaction(fn func(*Transaction) error) (err error) {
	var txn Transaction
	defer txn.close()

	if err = fn(&txn); err != nil {
		return
	}

	if err = a.logEntries(txn.entries); err != nil {
		return
	}

	// Ensure logs have flushed to disk
	return a.Flush()
}

// Close will close an instance of Actions
func (a *Actions) Close() (err error) {
	return a.Logger.Close()
}
