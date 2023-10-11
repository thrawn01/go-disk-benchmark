package benchmark

import "os"

type Writer interface {
	Write(data []byte) (int, error)
	Close() error
}

type blockWriter struct {
	file *os.File
}

func Block(filename string) Writer {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	return &blockWriter{file: f}
}

func (fw *blockWriter) Write(data []byte) (int, error) {
	return fw.file.Write(data)
}

func (fw *blockWriter) Close() error {
	return fw.file.Close()
}
