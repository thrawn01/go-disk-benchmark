package benchmark

import (
	"fmt"
	"github.com/akrylysov/pogreb"
	"github.com/thrawn01/go-disk-benchmark/uuid"
)

type pogrebWriter struct {
	uuid *uuid.Generator
	db   *pogreb.DB
}

func Pogreb(path string) *pogrebWriter {
	db, err := pogreb.Open(path, nil)
	if err != nil {
		panic(fmt.Errorf("failed to open Pogreb DB: %w", err))
	}
	return &pogrebWriter{
		uuid: uuid.MustNewGenerator(),
		db:   db,
	}
}

func (pw *pogrebWriter) Write(data []byte) (int, error) {
	key := pw.uuid.Next()

	if err := pw.db.Put(key[:], data); err != nil {
		return 0, err
	}
	return len(data), nil
}

func (pw *pogrebWriter) Close() error {
	return pw.db.Close()
}
