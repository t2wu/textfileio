package mybufio

import "os"

type osinterface interface {
	Create(name string) (*os.File, error)
	Open(name string) (*os.File, error)
}

// Create a file
func Create(name string) (*os.File, error) {
	return os.Create(name)
}

// Open a file
func Open(name string) (*os.File, error) {
	return os.Open(name)
}
