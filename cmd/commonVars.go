/*
Copyright © 2025 Route 1337 LLC.
This file is part of Fastbound Downloader.
*/

package cmd

import (
	"path/filepath"
	"runtime"
)

// The version string should be updated before any merge to main
var shortVersion = "0.2.1"
var projectMaintainer = "Route 1337 LLC"
var projectLicense = "MIT"
var functionHelpShort = "An automated way to keep compliant Fastbound A&D book downloads"
var functionHelpLong = `This tool is used to keep compliant automated downloads of
Fastbound A&D books locally using Docker/K8s vs the PowerShell script Fastbound provides.`
var buildArch = runtime.GOARCH
var SettingsFilePath, _ = filepath.Abs("/config/settings.json")
var fastboundAPIBaseURL = "https://cloud.fastbound.com"
