package benchmark

import (
	"fmt"
	"github.com/thrawn01/go-disk-benchmark/uuid"
	bolt "go.etcd.io/bbolt"
)

type boltWriter struct {
	uuid   *uuid.Generator
	db     *bolt.DB
	bucket []byte
}

func Bolt(path string) *boltWriter {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		panic(fmt.Errorf("failed to open BoltDB: %w", err))
	}

	// Ensure the bucket exists
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("bench"))
		return err
	})
	if err != nil {
		panic(err)
	}

	return &boltWriter{
		uuid:   uuid.MustNewGenerator(),
		bucket: []byte("bench"),
		db:     db,
	}
}

func (bw *boltWriter) Write(data []byte) (int, error) {
	key := bw.uuid.Next()

	err := bw.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bw.bucket)
		return b.Put(key[:], data)
	})

	if err != nil {
		return 0, err
	}

	return len(data), nil
}

func (bw *boltWriter) Close() error {
	return bw.db.Close()
}
