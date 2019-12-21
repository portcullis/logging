package logging

// Writer interface for putting log Entry somewhere, someway
type Writer interface {
	Write(e Entry)
}

// WriterFunc implements the logging.Writer interface to cast functions to the Writer
type WriterFunc func(e Entry)

func (w WriterFunc) Write(e Entry) {
	w(e)
}

// Discard will drop the log entries into the abyss (or do nothing)
var Discard = WriterFunc(func(Entry) {})
