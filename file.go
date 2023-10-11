package benchmark

import (
	"fmt"
	"os"
	"path/filepath"
)

type fileWriter struct {
	dir string
	seq int
}

func File(directory string) *fileWriter {
	if err := os.MkdirAll(directory, 0755); err != nil {
		panic(err)
	}

	// Initialize with the next available sequence number
	files, err := os.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	// We assume files in the directory are named as 1.dat, 2.dat, ...
	seq := len(files) + 1

	fw := &fileWriter{
		dir: directory,
		seq: seq,
	}

	return fw
}

func (fw *fileWriter) Write(data []byte) (int, error) {
	filename := filepath.Join(fw.dir, fmt.Sprintf("%d.dat", fw.seq))
	file, err := os.Create(filename)
	if err != nil {
		return 0, err
	}
	defer func() { _ = file.Close() }()
	fw.seq++
	return file.Write(data)
}

func (fw *fileWriter) Close() error {
	return nil
}
