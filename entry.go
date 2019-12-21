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

func newEntryPool() *entryPool {
	return &entryPool{
		pool: sync.Pool{
			New: func() interface{} {
				return &Entry{}
			},
		},
	}
}

func (ep *entryPool) Get() *Entry {
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
