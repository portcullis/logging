package logging_test

import (
	"testing"

	"github.com/portcullis/logging"
)

func Benchmark_Log_WithFields(b *testing.B) {
	l := logging.New(
		logging.WithWriter(logging.Discard),
		logging.WithFields(
			logging.Field{Name: "Test", Value: "Hello"},
		),
	).WithFields(
		logging.Field{Name: "Value", Value: false},
	)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l.Info("This was a mistake %q", true, false)
	}
}

func Benchmark_Log(b *testing.B) {
	l := logging.New(
		logging.WithWriter(logging.Discard),
		logging.WithFields(
			logging.Field{Name: "Test", Value: "Hello"},
		),
	)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l.Info("This was a mistake %q", false)
	}
}
