package actions

import (
	"time"

	"github.com/Hatch1fy/logger"
	"github.com/Hatch1fy/snapshotter"
)

// NewImporter will return a new instance of importer
func NewImporter(dir, name string, be snapshotter.Backend, interval time.Duration, fn Handler) (ip *Importer, err error) {
	var i Importer
	if i.Importer, err = logger.NewImporter(dir, name, be, interval, i.handleImport); err != nil {
		return
	}

	i.fn = fn
	ip = &i
	return
}

// Importer will import action logs
type Importer struct {
	*logger.Importer

	fn Handler
}

func (i *Importer) handleImport(ts time.Time, line []byte) (err error) {
	a, msg := parseLine(line)
	return i.fn(ts, a, msg)
}
