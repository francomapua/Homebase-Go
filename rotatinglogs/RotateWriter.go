package rotatinglogs

import (
	"os"
	"sync"
)

// RotateWriter : Struct, please don't use this
type RotateWriter struct {
	lock     sync.Mutex
	filename string
	fp       *os.File
}

// New : Constructor
func New(filename string) *RotateWriter {
	writer := &RotateWriter{filename: filename}
	err := writer.Rotate()
	if err != nil {
		return nil
	}
	return writer
}

func (writer *RotateWriter) Write(output []byte) (bytesWritten int, err error) {
	writer.lock.Lock()
	defer writer.lock.Unlock()
	bytesWritten, err = writer.fp.Write(output)
	return bytesWritten, err
}

func (writer *RotateWriter) Rotate() (err error) {
	writer.lock.Lock()
	defer writer.lock.Unlock()

	// Close Existing File If Open
	if writer.fp != nil {
		err = writer.fp.Close()
		writer.fp = nil
		if err != nil {
			return err
		}
	}

	// Rename Dest File if it Exists
	_, err = os.Stat(writer.filename)
	if err == nil {

	}
	// Create File
	return err
}
