package actions

import (
	"fmt"
	"testing"
	"time"
)

func TestReader(t *testing.T) {
	var (
		a   *Actions
		err error
	)

	if err = initTestDirs(); err != nil {
		t.Fatal(err)
	}
	defer removeTestDirs()

	if a, err = New("./test_logs", "tester"); err != nil {
		t.Fatal(err)
	}
	defer a.Close()

	a.SetNumLines(5)

	if err = testPopulateActions(a, ActionCreate, 5, 0); err != nil {
		t.Fatal(err)
	}

	if err = testPopulateActions(a, ActionEdit, 5, 5); err != nil {
		t.Fatal(err)
	}

	if err = testPopulateActions(a, ActionDelete, 5, 10); err != nil {
		t.Fatal(err)
	}
}

func testReaderIteration(ts time.Time, a Action, key, value []byte, count int) (err error) {
	newKey := fmt.Sprintf("%d", count)
	newVal := fmt.Sprintf("#%d", count)

	if string(key) != newKey {
		return fmt.Errorf("invalid message, expected \"%s\" and received \"%s\"", newKey, string(key))
	}

	if string(value) != newVal {
		return fmt.Errorf("invalid message, expected \"%s\" and received \"%s\"", newVal, string(value))
	}

	switch count {
	case 1, 2, 3, 4, 5:
		if a != ActionCreate {
			return fmt.Errorf("invalid action, expected \"%s\" and received \"%s\"", ActionCreateStr, a.String())
		}

	case 6, 7, 8, 9, 10:
		if a != ActionEdit {
			return fmt.Errorf("invalid action, expected \"%s\" and received \"%s\"", ActionEditStr, a.String())
		}

	case 11, 12, 13, 14, 15:
		if a != ActionDelete {
			return fmt.Errorf("invalid action, expected \"%s\" and received \"%s\"", ActionDeleteStr, a.String())
		}

	}

	return
}
