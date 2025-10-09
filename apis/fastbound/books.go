/*
Copyright © 2025 Route 1337 LLC.
This file is part of Fastbound Downloader.
*/

package fastbound

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/route1337/fastbound-downloader/apis/fbdownloader_settings"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

// DownloadBoundBook downloads the latest A&D book from the Fastbound API and return the path of the saved file
func DownloadBoundBook(apiBase string, config fbdownloader_settings.FBDConfig) (string, error) {

	// Define a struct to hold the API's JSON response.
	type downloadApiResponse struct {
		URL string `json:"url"`
	}

	// Craft the API URL and set up a Context
	apiURL := fmt.Sprintf("%s/%s/api/Downloads/BoundBook", apiBase, config.Fastbound.AccountNumber)
	apiContext, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	client := &http.Client{}

	// Create a new HTTP request with context.
	postRequest, err := http.NewRequestWithContext(apiContext, "POST", apiURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create API POST request: %w", err)
	}

	postRequest.SetBasicAuth(config.Fastbound.ApiKey, config.Fastbound.ApiKey)
	postRequest.Header.Set("accept", "application/json")
	postRequest.Header.Set("X-AuditUser", config.Fastbound.AuditUser)

	// Execute the request using a default HTTP client.
	postResponse, err := client.Do(postRequest)
	if err != nil {
		return "", fmt.Errorf("failed to execute POST request: %w", err)
	}
	defer func() {
		if err := postResponse.Body.Close(); err != nil {
			log.Printf("Warning: failed to close postResponse body: %v", err)
		}
	}()

	// Read the response status code and fail out with any errors
	if postResponse.StatusCode != http.StatusOK {
		errorBody, _ := io.ReadAll(postResponse.Body)
		return "", fmt.Errorf("api request failed with status %d: %s", postResponse.StatusCode, string(errorBody))
	}
	// Decode the JSON response from the POST request
	var apiResponse downloadApiResponse
	if err := json.NewDecoder(postResponse.Body).Decode(&apiResponse); err != nil {
		return "", fmt.Errorf("failed to decode JSON response: %w", err)
	}
	if apiResponse.URL == "" {
		return "", fmt.Errorf("API response did not contain a download URL")
	}
	// Download the file from the provided URL
	downloadRequest, err := http.NewRequestWithContext(apiContext, "GET", apiResponse.URL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create GET request for download: %w", err)
	}
	downloadResponse, err := client.Do(downloadRequest)
	if err != nil {
		return "", fmt.Errorf("failed to download file from URL: %w", err)
	}
	defer func() {
		if err := downloadResponse.Body.Close(); err != nil {
			log.Printf("Warning: failed to close downloadResponse body: %v", err)
		}
	}()

	if downloadResponse.StatusCode != http.StatusOK {
		return "", fmt.Errorf("file download failed with status %d", downloadResponse.StatusCode)
	}
	// Extract file name from URL
	parsedUrl, err := url.Parse(apiResponse.URL)
	if err != nil {
		return "", fmt.Errorf("failed to parse download URL: %w", err)
	}
	downloadedBook := filepath.Base(parsedUrl.Path)
	// Set a destination path and create a file to store there
	destinationPath := filepath.Join(config.Paths.BoundBooks, downloadedBook)
	storeFile, err := os.Create(destinationPath)
	if err != nil {
		return "", fmt.Errorf("failed to save bound book file: %w", err)
	}
	defer func() {
		if err := storeFile.Close(); err != nil {
			log.Printf("Warning: failed to close storeFile: %v", err)
		}
	}()

	// Stream the file contents to the new file
	if _, err := io.Copy(storeFile, downloadResponse.Body); err != nil {
		return "", fmt.Errorf("failed to write the bound book file: %w", err)
	}

	return destinationPath, nil
}
