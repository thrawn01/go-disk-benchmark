package uuid_test

import (
	"github.com/thrawn01/go-disk-benchmark/uuid"
	"testing"
)

var _s string

func BenchmarkHex128(b *testing.B) {
	g := uuid.MustNewGenerator()
	for i := 0; i < b.N; i++ {
		_s = g.Hex128()
	}
}

func BenchmarkNext(b *testing.B) {
	g := uuid.MustNewGenerator()
	for i := 0; i < b.N; i++ {
		g.Next()
	}
}
