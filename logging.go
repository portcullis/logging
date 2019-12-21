package logging

import (
	"strings"
	"sync"
	"time"
)

// Log for ... writing logs
type Log struct {
	entries *entryPool
	once    sync.Once
	parent  *Log
	fields  []Field
	writer  Writer
}

// New creates a new Log instance with the specified options
func New(opts ...Option) *Log {
	l := &Log{}
	l.once.Do(l.initialize)

	for _, opt := range opts {
		opt(l)
	}

	return l
}

// WithFields will create a child logger sharing the same output as the parent with the additional fields specified
func (l *Log) WithFields(fields ...Field) *Log {
	child := &Log{
		parent: l,
		fields: fields,
	}

	l.once.Do(l.initialize)

	return child
}

func (l *Log) initialize() {
	l.entries = newEntryPool()
	l.fields = make([]Field, 0)
	l.writer = Discard
}

func (l *Log) Write(lvl Level, msg string, args ...interface{}) {
	l.once.Do(l.initialize)

	if l.parent != nil {
		l.parent.innerWriter(lvl, msg, args, append(l.parent.fields, l.fields...))
	} else {
		l.innerWriter(lvl, msg, args, l.fields)
	}
}

// Critical will write a Critical log entry
func (l *Log) Critical(msg string, args ...interface{}) {
	l.Write(LevelCritical, msg, args...)
}

// Error will write an Error log entry
func (l *Log) Error(msg string, args ...interface{}) {
	l.Write(LevelError, msg, args...)
}

// Warning will write a Warning log entry
func (l *Log) Warning(msg string, args ...interface{}) {
	l.Write(LevelWarning, msg, args...)
}

// Info will write an Informational log entry
func (l *Log) Info(msg string, args ...interface{}) {
	l.Write(LevelInformational, msg, args...)
}

// Debug will write a Debug log entry
func (l *Log) Debug(msg string, args ...interface{}) {
	l.Write(LevelDebug, msg, args...)
}

// Trace will write a Trace log entry
func (l *Log) Trace(msg string, args ...interface{}) {
	l.Write(LevelTrace, msg, args...)
}

func (l *Log) innerWriter(lvl Level, msg string, args []interface{}, fields []Field) {
	l.once.Do(l.initialize)

	if l.writer == nil {
		return
	}

	e := l.entries.Get()
	e.Timestamp = time.Now()
	e.Level = lvl
	e.Message = strings.TrimSpace(msg)

	for i := range args {
		e.Arguments = append(e.Arguments, args[i])
	}

	e.Fields = l.fields[:]

	defer l.entries.Put(e)

	l.writer.Write(*e)
}
