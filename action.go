package actions

const (
	// ActionNil represents an unset action
	ActionNil Action = iota
	// ActionCreate represents a creation action
	ActionCreate
	// ActionEdit represents an editing action
	ActionEdit
	// ActionDelete represents a deleting action
	ActionDelete
)

const (
	// ActionNilStr is a string representation of ActionNil
	ActionNilStr = ""
	// ActionCreateStr is a string representation of ActionCreate
	ActionCreateStr = "CREATE"
	// ActionEditStr is a string representation of ActionEdit
	ActionEditStr = "EDIT"
	// ActionDeleteStr is a string representation of ActionDelete
	ActionDeleteStr = "DELETE"
)

// ParseAction will parse a stringified action and return it as an action
func ParseAction(action string) Action {
	switch action {
	case ActionCreateStr:
		return ActionCreate
	case ActionEditStr:
		return ActionEdit
	case ActionDeleteStr:
		return ActionDelete
	}

	return ActionNil
}

// Action represents a service action
type Action uint8

func (a Action) String() string {
	switch a {
	case ActionCreate:
		return ActionCreateStr
	case ActionEdit:
		return ActionEditStr
	case ActionDelete:
		return ActionDeleteStr
	}

	return ActionNilStr
}

// MarshalJSON is a JSON marshaling helper func
func (a Action) MarshalJSON() (b []byte, err error) {
	return
}
