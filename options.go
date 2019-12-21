package logging

// Option for logging
type Option func(*Log)

// Apply aa Option to the specified Log
func (o Option) Apply(l *Log) {
	if o == nil {
		return
	}

	o(l)
}

// WithFields options will append the specified fields to the Log
func WithFields(fields ...Field) Option {
	return func(l *Log) {
		for _, f := range fields {
			l.fields = append(l.fields, f)
		}
	}
}

// WithWriter option will set the underlying write for the logs
func WithWriter(w Writer) Option {
	return func(l *Log) {
		if w == nil {
			l.writer = Discard
		} else {
			l.writer = w
		}
	}
}
