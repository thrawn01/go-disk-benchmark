package benchmark

import (
	"context"
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/thrawn01/go-disk-benchmark/uuid"
)

type PebbleWriter struct {
	uuid *uuid.Generator
	db   *pebble.DB
}

func Pebble(path string) *PebbleWriter {
	db, err := pebble.Open(path, &pebble.Options{
		Logger: NoopLoggerAndTracer{},
	})
	if err != nil {
		panic(fmt.Errorf("failed to open Pebble DB: %w", err))
	}
	return &PebbleWriter{
		uuid: uuid.MustNewGenerator(),
		db:   db,
	}
}

func (pw *PebbleWriter) Write(data []byte) (int, error) {
	key := pw.uuid.Next()

	if err := pw.db.Set(key[:], data, pebble.Sync); err != nil {
		return 0, err
	}

	return len(data), nil
}

func (pw *PebbleWriter) Close() error {
	return pw.db.Close()
}

// NoopLoggerAndTracer does no logging and tracing. Remember that struct{} is
// special cased in Go and does not incur an allocation when it backs the
// interface LoggerAndTracer.
type NoopLoggerAndTracer struct{}

// Infof implements LoggerAndTracer.
func (l NoopLoggerAndTracer) Infof(format string, args ...interface{}) {}

// Fatalf implements LoggerAndTracer.
func (l NoopLoggerAndTracer) Fatalf(format string, args ...interface{}) {}

// Eventf implements LoggerAndTracer.
func (l NoopLoggerAndTracer) Eventf(ctx context.Context, format string, args ...interface{}) {
}
