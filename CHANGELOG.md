Changelog
---------
A list of changes made to Fastbound Downloader

Version 0.2.1
-------------

1. Define scanning interval in minutes with a 1 day default setting if not defined
2. Start with a forced 5 minute delay if metrics are enabled before first scan

Version 0.2.0
-------------

1. Detect if the download has already been processed. If so skip downloading again and log the skip.
2. Add the following metrics for Prometheus or other observability tooling
    1. `DownloadedBooksTotal` - A total count of downloaded bound books
    2. `SkippedBookDownloadsTotal` - A total count of bound books downloads that were skipped due to an existing file
    3. `FailedBookDownloadsTotal` - A total count of failures to download a bound book

Version 0.1.0
-------------

1. Initial release
