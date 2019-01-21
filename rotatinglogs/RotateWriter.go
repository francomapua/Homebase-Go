package rotatinglogs

import (
	"os"
	"sync"
	"time"
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

// Write : writes a log to a file
func (writer *RotateWriter) Write(output []byte) (bytesWritten int, err error) {
	writer.lock.Lock()
	defer writer.lock.Unlock()
	bytesWritten, err = writer.fp.Write(output)
	return bytesWritten, err
}

// Rotate : archives and deletes files after a certain number of rotations
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
	if _, err = os.Stat(writer.filename); !os.IsNotExist(err) {
		err = os.Rename(writer.filename, writer.filename+"."+time.Now().Format(time.RFC3339))
		if err != nil {
			return err
		}
	}

	// Delete the old files
	// Create File
	writer.fp, err = os.Create(writer.filename)
	return err
}
