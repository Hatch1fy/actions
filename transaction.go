package actions

import "encoding/json"

// Transaction allows a batch of logs to occur without flushing
type Transaction struct {
	entries []transactionEntry
}

func (t *Transaction) close() {
	t.entries = nil
}

// Log will log an action with a byteslice message
func (t *Transaction) Log(action Action, key, value []byte) (err error) {
	entry := newTransactionEntry(action, key, value)
	t.entries = append(t.entries, entry)
	return
}

// LogString will log an action with a string message
func (t *Transaction) LogString(action Action, key, value string) (err error) {
	return t.Log(action, []byte(key), []byte(value))
}

// LogJSON will log an action with a JSON message
func (t *Transaction) LogJSON(action Action, key []byte, value interface{}) (err error) {
	var bs []byte
	if bs, err = json.Marshal(value); err != nil {
		return
	}

	return t.Log(action, key, bs)
}

func newTransactionEntry(action Action, key, value []byte) (e transactionEntry) {
	e.action = action
	e.key = key
	e.value = value
	return
}

type transactionEntry struct {
	action Action
	key    []byte
	value  []byte
}
