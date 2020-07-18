package logging

import (
	"sync"
	"time"
)

// Field for structured logging
type Field struct {
	Name  string
	Value interface{}
}

// Entry represents a single log
type Entry struct {
	Level     Level
	Timestamp time.Time
	Message   string
	Arguments []interface{}
	Fields    []Field
}

type entryPool struct {
	pool sync.Pool
}

var entryPoolConstructor = func() interface{} { return &Entry{} }

func (ep *entryPool) Get() *Entry {
	// set every time, so we know it is set, most effecient to keep locks and initializers out of the way
	ep.pool.New = entryPoolConstructor
	return ep.pool.Get().(*Entry)
}

func (ep *entryPool) Put(e *Entry) {
	e.Level = 0
	e.Message = ""
	e.Timestamp = time.Time{}
	e.Arguments = e.Arguments[:0]
	e.Fields = e.Fields[:0]

	ep.pool.Put(e)
}
