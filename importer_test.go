package actions

import (
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/Hatch1fy/logger"
	"github.com/Hatch1fy/snapshotter/backends"
)

func TestImporter(t *testing.T) {
	var (
		a   *Actions
		i   *Importer
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

	be := backends.NewFile("./test_backend")
	ss := logger.NewSnapshotter(be, true)

	a.SetNumLines(5)
	a.SetRotateFn(func(filename string) {
		if err = ss.Snapshot(filename); err != nil {
			t.Fatal(err)
		}
	})

	var count int
	if i, err = NewImporter("./test_data", "tester", be, time.Hour, func(ts time.Time, a Action, msg []byte) (err error) {
		count++
		newMsg := fmt.Sprintf("#%d", count)
		if string(msg) != newMsg {
			return fmt.Errorf("invalid message, expected \"%s\" and received \"%s\"", newMsg, string(msg))
		}

		return
	}); err != nil {
		return
	}
	defer i.Close()

	if err = testPopulateActions(a, ActionCreate, 5, 0); err != nil {
		t.Fatal(err)
	}

	if err = testPopulateActions(a, ActionEdit, 5, 5); err != nil {
		t.Fatal(err)
	}

	if err = testPopulateActions(a, ActionDelete, 5, 10); err != nil {
		t.Fatal(err)
	}

	if err = i.Import(); err != nil {
		t.Fatal(err)
	}

	if err = i.Import(); err != nil {
		t.Fatal(err)
	}

	if err = i.Import(); err != nil {
		t.Fatal(err)
	}

	if err = i.Import(); err != io.EOF {
		t.Fatal(err)
	}
}
