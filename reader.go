package actions

import (
	"time"

	"github.com/Hatch1fy/logger"
)

// NewReader will return a new instance of reader
func NewReader(filename string) (rp *Reader, err error) {
	var r Reader
	if r.r, err = logger.NewReader(filename); err != nil {
		return
	}

	rp = &r
	return
}

// Reader will read action logs
type Reader struct {
	r *logger.Reader
}

// ForEach will iterate through each line, starting at the offset
func (r *Reader) ForEach(offset int64, fn Handler) (err error) {
	err = r.r.ForEach(offset, func(ts time.Time, line []byte) (err error) {
		a, key, value := parseLine(line)
		return fn(ts, a, key, value)
	})

	return
}

// Close will close an instance of reader
func (r *Reader) Close() (err error) {
	return r.r.Close()
}
