/*
Copyright © 2025 Route 1337 LLC.
This file is part of Fastbound Downloader.
*/

package fastbound

import (
	"fmt"
	"github.com/route1337/fastbound-downloader/apis/fbdownloader_settings"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestDownloadBoundBook validates the DownloadBoundBook function
func TestDownloadBoundBook(t *testing.T) {
	// Create a mock server to simulate the Fastbound API
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a mock API POST request and response
		if r.Method == "POST" && strings.Contains(r.URL.Path, "/api/Downloads/BoundBook") {
			// Force overriding the returned URL
			responseURL := "http://" + r.Host + "/download/MOCK_BOUND_BOOK.pdf"
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			// Return what would be a valid result for the API
			_, err := fmt.Fprintf(w, `{"url": "%s"}`, responseURL)
			if err != nil {
				t.Fatalf("Mock server failed to write response: %v", err)
			}
			return
		}
		// Request the download of the mocked bound book
		if r.Method == "GET" && r.URL.Path == "/download/MOCK_BOUND_BOOK.pdf" {
			w.WriteHeader(http.StatusOK)
			// We're writing some dummy data here
			_, err := w.Write([]byte(`"Guns. Lots of guns."`))
			if err != nil {
				t.Fatalf("Mock server failed to write file content: %v", err)
			}
			return
		}
		// Return a 404 if the request doesn't match any of the above
		t.Errorf("Mock server received unexpected request: %s %s", r.Method, r.URL.Path)
		http.NotFound(w, r)
	}))
	defer mockServer.Close()

	tempDir := t.TempDir()
	// Create a test config for use with function calls
	testConfig := fbdownloader_settings.FBDConfig{
		Fastbound: struct {
			AccountNumber string `json:"account-number"`
			ApiKey        string `json:"api-key"`
			AuditUser     string `json:"audit-user"`
		}{
			AccountNumber: "123456",
			ApiKey:        "kkJ4K3dHoHqZzNvoDJ",
			AuditUser:     "pgibbons@initech.com",
		},
		Paths: struct {
			BoundBooks       string `json:"bound-books"`
			BackgroundChecks string `json:"background-checks"`
		}{
			BoundBooks:       tempDir,
			BackgroundChecks: tempDir,
		},
	}

	// Call the DownloadBoundBook function using our mockServer URL instead of the real API
	savedFilePath, err := DownloadBoundBook(mockServer.URL, testConfig)
	if err != nil {
		t.Fatalf("DownloadBoundBook() returned an unexpected error: %v", err)
	}

	// Check that the returned path is correct
	expectedFile := filepath.Join(tempDir, "MOCK_BOUND_BOOK.pdf")
	if savedFilePath != expectedFile {
		t.Errorf("Expected saved file path to be '%s', but got '%s'", expectedFile, savedFilePath)
	}

	// Check that the file was actually created on disk
	if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
		t.Errorf("Expected file to be created, but it was not: %s", expectedFile)
	}
}
