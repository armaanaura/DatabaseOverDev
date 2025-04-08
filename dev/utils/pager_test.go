package utils

import (
	"os"
	"testing"
	"encoding/binary"
)

func TestCreatePage(t *testing.T) {
	tempFile := "test_database.db"
	// defer os.Remove(tempFile) // clean up after test

	// Create an empty temp DB file
	// file, err := os.Create(tempFile)
	// if err != nil {
	// 	t.Fatalf("Failed to create temp file: %v", err)
	// }
	// file.Close()

	pageNumber := 1
	pageType := "Table Page"
	pageSize := 4096

	err := CreatePage(tempFile, pageNumber, pageType, pageSize)
	if err != nil {
		t.Fatalf("CreatePage failed: %v", err)
	}

	// Reopen file to read contents
	data, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("Failed to read database file: %v", err)
	}

	// Validate page size
	if len(data) != pageSize {
		t.Fatalf("Expected page size %d, got %d", pageSize, len(data))
	}

	// Validate page number written in first 4 bytes
	pageNum := binary.LittleEndian.Uint32(data[0:4])
	if pageNum != uint32(pageNumber) {
		t.Errorf("Expected page number %d, got %d", pageNumber, pageNum)
	}

	// Validate page type at byte 4
	if data[4] != 1 { // "Meta Page" = 0
		t.Errorf("Expected page type 0, got %d", data[4])
	}
}
