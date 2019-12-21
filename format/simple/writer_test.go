package simple_test

import (
	"io/ioutil"
	"testing"

	"github.com/portcullis/logging"
	"github.com/portcullis/logging/format/simple"
)

func BenchmarkSimple_Write(b *testing.B) {
	log := logging.New(
		logging.WithWriter(simple.New(ioutil.Discard, simple.Level(logging.LevelInformational))),
		logging.WithFields(
			logging.Field{Name: "app", Value: "sandbox"},
		),
	)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		log.Info("Test")
	}
}
