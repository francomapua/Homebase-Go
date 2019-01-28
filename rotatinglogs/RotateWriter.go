// Version 2
package logger

import (
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	maxLogFiles int = 5
)

// RotateWriter : Struct
type RotateWriter struct {
	lock     sync.Mutex
	filename string
	fp       *os.File
}

// NewRotateWriter : Constructor
func NewRotateWriter(filename string) *RotateWriter {
	filename, _ = filepath.Abs(filename)
	writer := &RotateWriter{filename: filename}
	return writer
}

// Write : Basic logging function
func (writer *RotateWriter) Write(output []byte) (bytesWritten int, err error) {
	// Check for Rotation
	if stat, err := os.Stat(writer.filename); os.IsNotExist(err) || !writer.areDatesOnSameDay(stat.ModTime(), time.Now()) {
		// If log file does not exist OR log file was last editted before today
		writer.rotate()
	}

	// Mutex Lock
	writer.lock.Lock()
	defer writer.lock.Unlock()

	// Write to file
	writer.fp, err = os.OpenFile(writer.filename, os.O_WRONLY|os.O_APPEND, 0644)
	defer writer.fp.Close()
	bytesWritten, err = writer.fp.Write(output)

	return bytesWritten, err
}

func (writer *RotateWriter) rotate() (err error) {
	// Mutex Lock
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

	// Rename Old File If it Exists
	if stat, err := os.Stat(writer.filename); !os.IsNotExist(err) {
		err = os.Rename(writer.filename, writer.filename+"-"+writer.getSimpleDate(stat.ModTime()))
		if err != nil {
			return err
		}
	}

	// Create File
	writer.upsertDirectory(filepath.Dir(writer.filename))
	writer.fp, err = os.Create(writer.filename)

	// Trim files if there are too many logs
	var logFiles []os.FileInfo
	dir, _ := filepath.Abs(filepath.Dir(writer.filename))
	logname := filepath.Base(writer.filename)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		name := info.Name()
		print(name)
		if strings.Contains(info.Name(), logname) {
			logFiles = append(logFiles, info)
		}
		return nil
	})
	if len(logFiles) > maxLogFiles {
		// Sort Slice by edit date
		sort.Slice(logFiles, func(a, b int) bool {
			return logFiles[a].ModTime().After(logFiles[b].ModTime())
		})
		// Delete Other Logs
		for i := len(logFiles); i > maxLogFiles; i-- {
			deletePath, _ := filepath.Abs(dir + "/" + logFiles[i-1].Name())
			os.Remove(deletePath)
		}
	}
	return nil
}

/* Helper Functions */
func (writer *RotateWriter) areDatesOnSameDay(dateA time.Time, dateB time.Time) bool {
	yearA, monthA, dayA := dateA.Date()
	yearB, monthB, dayB := dateB.Date()
	return yearA == yearB && monthA == monthB && dayA == dayB
}
func (writer *RotateWriter) getSimpleDate(dateTime time.Time) string {
	year, month, day := dateTime.Date()
	runes := []rune(strings.ToUpper(month.String()))
	monthShort := string(runes[0:3])
	concatString := strconv.Itoa(year) + monthShort + strconv.Itoa(day)
	return concatString
}
func (writer *RotateWriter) upsertDirectory(path string) error {
	var err error
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
	}
	return err
}
