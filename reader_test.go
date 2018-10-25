package actions

import (
	"fmt"
	"path"
	"testing"
	"time"

	"github.com/Hatch1fy/snapshotter/backends"
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

	var count int
	be := backends.NewFile("./test_logs")
	if err = be.ForEach("tester", "", 3, func(filename string) (err error) {
		var r *Reader
		if r, err = NewReader(path.Join("./test_logs", filename)); err != nil {
			return
		}
		defer r.Close()

		return r.ForEach(0, func(ts time.Time, a Action, msg []byte) (err error) {
			count++
			return testReaderIteration(ts, a, msg, count)
		})
	}); err != nil {
		t.Fatal(err)
	}

	if count != 15 {
		t.Fatalf("invalid count, expected %d and received %d", 15, count)
	}
}

func testReaderIteration(ts time.Time, a Action, msg []byte, count int) (err error) {
	newMsg := fmt.Sprintf("#%d", count)
	if string(msg) != newMsg {
		return fmt.Errorf("invalid message, expected \"%s\" and received \"%s\"", string(msg), newMsg)
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
