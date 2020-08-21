package actions

import (
	"fmt"
	"os"
	"testing"

	"github.com/hatchify/errors"
)

func TestActions(t *testing.T) {
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

	if err = testPopulateActions(a, ActionCreate, 15, 0); err != nil {
		t.Fatal(err)
	}
}

func testPopulateActions(a *Actions, action Action, count, offset int) (err error) {
	for i := 0; i < count; i++ {
		key := fmt.Sprintf("%d", i+1+offset)
		value := fmt.Sprintf("#%d", i+1+offset)
		if err = a.LogString(action, key, value); err != nil {
			return
		}
	}

	return
}

func initTestDirs() (err error) {
	if err = os.MkdirAll("./test_logs", 0744); err != nil {
		return
	}

	if err = os.MkdirAll("./test_backend", 0744); err != nil {
		return
	}

	if err = os.MkdirAll("./test_data", 0744); err != nil {
		return
	}

	return
}

func removeTestDirs() (err error) {
	var errs errors.ErrorList
	errs.Push(os.RemoveAll("./test_logs"))
	errs.Push(os.RemoveAll("./test_backend"))
	errs.Push(os.RemoveAll("./test_data"))
	return errs.Err()
}
