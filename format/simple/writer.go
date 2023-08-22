// Deprecated: This package has been deprecated and log/slog is preferred
package simple

import (
	"fmt"
	"io"
	"strings"

	"github.com/portcullis/logging"
)

// Option for the simple Writer
type Option func(w *Writer)

// Writer is a simple logging.Writer implementation
type Writer struct {
	w          io.Writer
	lvl        logging.Level
	timeFormat string
}

// New creates a new Writer to output the specified io.Writer with the provided Options
//
// This is a super simple implementation (hence the name) that will output in the following format:
// > [<timestamp>] <level> <message> <field.Key=field.Value>...
func New(w io.Writer, opts ...Option) *Writer {
	writer := &Writer{
		w:          w,
		timeFormat: "2006-01-02 15:04:05",
		lvl:        logging.LevelInformational,
	}

	for _, opt := range opts {
		opt(writer)
	}

	return writer
}

// Level will apply the logging level to the writer for output verbosity
func Level(lvl logging.Level) Option {
	return func(w *Writer) {
		w.lvl = lvl
	}
}

func (w *Writer) Write(e logging.Entry) {
	if !e.Level.Is(w.lvl) {
		return
	}

	// TODO: Make this have as close to zero allcoations as possible
	// current implementation is just to get it working with the target formatting
	fields := ""
	for _, field := range e.Fields {
		fields += " " + field.Name + "=" + fmt.Sprintf("%v", field.Value)
	}

	fmt.Fprintln(w.w, "["+e.Timestamp.Format(w.timeFormat)+"] "+strings.ToUpper(e.Level.Short())+" "+fmt.Sprintf(e.Message, e.Arguments...)+fields)
}
