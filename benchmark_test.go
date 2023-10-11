package benchmark_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thrawn01/go-disk-benchmark"
)

const MB = 1024 * 1024

var conf = benchmark.CassandraConfig{
	Hosts:    []string{"127.0.0.1"},
	Keyspace: "benchmark",
	Table:    "benchmark",
	// Add these when running in production or staging environment
	DataCenter: "",
	Rack:       "",
}

func generateData(size int) []byte {
	return bytes.Repeat([]byte("A"), size)
}

func benchmarkWriter(writer benchmark.Writer, dataSize int, b *testing.B) {
	data := generateData(dataSize)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := writer.Write(data)
		if err != nil {
			b.Fatal(err)
		}
	}
	_ = writer.Close()
}

func TestCreateCassandraTable(t *testing.T) {
	require.NoError(t, benchmark.CreateTable(conf))
}

func BenchmarkBlock_100KB(b *testing.B) {
	benchmarkWriter(benchmark.Block("data/100kb.block"), 100*1024, b)
}

func BenchmarkPog_100KB(b *testing.B) {
	benchmarkWriter(benchmark.Pogreb("data/100kb.pog"), 100*1014, b)
}

func BenchmarkFile_100KB(b *testing.B) {
	benchmarkWriter(benchmark.File("data/100kb.file"), 100*1024, b)
}

func BenchmarkCassandra_100KB(b *testing.B) {
	benchmarkWriter(benchmark.Cassandra(conf), 100*1024, b)
}

func BenchmarkPebble_100KB(b *testing.B) {
	benchmarkWriter(benchmark.Pebble("data/100kb.pebble"), 100*1024, b)
}

func BenchmarkBolt_100KB(b *testing.B) {
	benchmarkWriter(benchmark.Bolt("data/100kb.bolt"), 100*1024, b)
}

// -------------

func BenchmarkBlock_10MB(b *testing.B) {
	benchmarkWriter(benchmark.Block("data/10mb.block"), 10*MB, b)
}

func BenchmarkFile_10MB(b *testing.B) {
	benchmarkWriter(benchmark.File("data/10mb.file"), 10*MB, b)
}

func BenchmarkPog_10MB(b *testing.B) {
	benchmarkWriter(benchmark.Pogreb("data/10mb.pog"), 10*MB, b)
}

func BenchmarkCassandra_10MB(b *testing.B) {
	benchmarkWriter(benchmark.Cassandra(conf), 10*MB, b)
}

func BenchmarkPebble_10MB(b *testing.B) {
	benchmarkWriter(benchmark.Pebble("data/10mb.pebble"), 10*MB, b)
}

func BenchmarkBolt_10MB(b *testing.B) {
	benchmarkWriter(benchmark.Bolt("data/10mb.bolt"), 10*MB, b)
}

// ---------------

func BenchmarkBlock_25MB(b *testing.B) {
	benchmarkWriter(benchmark.Block("data/25mb.block"), 25*MB, b)
}

func BenchmarkFile_25MB(b *testing.B) {
	benchmarkWriter(benchmark.File("data/25mb.file"), 25*MB, b)
}

func BenchmarkPog_25MB(b *testing.B) {
	benchmarkWriter(benchmark.Pogreb("data/25mb.pog"), 25*MB, b)
}

func BenchmarkPebble_25MB(b *testing.B) {
	benchmarkWriter(benchmark.Pebble("data/25mb.pebble"), 25*MB, b)
}

func BenchmarkBolt_25MB(b *testing.B) {
	benchmarkWriter(benchmark.Bolt("data/25mb.bolt"), 25*MB, b)
}

//func BenchmarkWriterCassandra25mb(b *testing.B) {
//	benchmarkWriter(benchmark.Cassandra(conf), 25*MB, b)
//}
