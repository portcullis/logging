package logging_test

import (
	"testing"

	"github.com/portcullis/logging"
)

func Benchmark_Log(b *testing.B) {
	l := logging.New(
		logging.WithFields(
			logging.Field{Name: "Test", Value: "Hello"},
		),
		logging.WithWriter(logging.Discard),
	).WithFields(
		logging.Field{Name: "Value", Value: false},
	)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l.Info("This was a mistake %q", true)
	}
}
