/*
Copyright © 2025 Route 1337 LLC.
This file is part of Fastbound Downloader.
*/

package cmd

import (
	"github.com/route1337/fastbound-downloader/apis/fastbound"
	"github.com/route1337/fastbound-downloader/apis/fbdownloader_settings"
	"log"
)

// rotationCycle This function runs the core logic of the Fastbound Downloader
func rotationCycle(settings fbdownloader_settings.FBDConfig) {
	log.Printf("Downloading the latest bound book for account %s\n", settings.Fastbound.AccountNumber)
	// Download the daily Bound Book
	downloadedBook, err := fastbound.DownloadBoundBook(fastboundAPIBaseURL, settings)
	if err != nil {
		return
	}
	log.Printf("Downloaded the bound book %s\n", downloadedBook)
}
