package io

import (
	"io"
	"io/ioutil"
	"os"
)

// Exists : Synchronously checks if file exists
func Exists(fPath string) bool {
	_, err := os.Stat(fPath)
	return !os.IsNotExist(err) // err would be not-exists
}

// ReadFileAsByte : Reads File in its entirety and Returns as Byte Array
func ReadFileAsByte(path string) ([]byte, error) {
	byteArr, err := ioutil.ReadFile(path)
	return byteArr, err
}

// CopySmallFile : copies small file by reading and writing it in its entirey
func CopySmallFile(inputPath, outputPath string) error {
	byteArr, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputPath, byteArr, 0644)
	return err
}

// CopyLargeFile : copies a large file through a cache
func CopyLargeFile(inputPath, outputPath string) error {
	// Get Input File
	ptrInputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer ptrInputFile.Close()

	// Get Output File
	ptrOutputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer ptrOutputFile.Close()

	// Create Buffer to keep read chunks
	buffer := make([]byte, 1024)
	for {
		// Read Chunk
		nBytesRead, err := ptrInputFile.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		} else if nBytesRead == 0 {
			break
		}
		// Write a Chunk
		if _, err := ptrOutputFile.Write(buffer[:nBytesRead]); err != nil {
			return err
		}
	}
	return nil
}

// UpsertDirectory : creates a directory if it exists
func UpsertDirectory(path string) error {
	var err error
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
	}
	return err
}

//func UpsertFile()
