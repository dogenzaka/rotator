package rotator

// Rotator interface
type Rotator interface {
	// Write binaries to the file.
	Write(bytes []byte) (n int, err error)
	// WriteString writes strings to the file.
	WriteString(str string) (n int, err error)
	// Close the file
	Close() error
}
