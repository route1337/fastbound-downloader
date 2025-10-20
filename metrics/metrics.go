/*
Copyright © 2025 Route 1337 LLC.
This file is part of Fastbound Downloader.
*/

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// DownloadedBooksTotal counts the total number of successful bound book downloads
	DownloadedBooksTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "fastbound_downloader_downloaded_books_total",
		Help: "The total number of successful bound book downloads",
	})

	// SkippedBookDownloadsTotal counts the total number of bound books downloads that were skipped due to an existing file
	SkippedBookDownloadsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "fastbound_downloader_skipped_book_downloads_total",
		Help: "The total number of times the found book was already detected as downloaded",
	})

	// FailedBookDownloadsTotal counts the total number of failed bound book downloads
	FailedBookDownloadsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "fastbound_downloader_failed_book_downloads_total",
		Help: "The total number of failed attempts at downloading a bound book",
	})
)
