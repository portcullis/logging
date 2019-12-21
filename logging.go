package logging

import (
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

func (l *Log) innerWriter(lvl Level, msg string, args []interface{}, fields []Field) {
	l.once.Do(l.initialize)

	if l.writer == nil {
		return
	}

	e := l.entries.Get()
	e.Timestamp = time.Now()
	e.Level = lvl
	e.Message = msg

	for i := range args {
		e.Arguments = append(e.Arguments, args[i])
	}

	e.Fields = l.fields[:]

	defer l.entries.Put(e)

	l.writer.Write(*e)
}
