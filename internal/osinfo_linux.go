// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016 Datadog, Inc.

package internal

import (
	"bufio"
	"os"
	"strings"
)

const (
	unknown = "unknown"
)

// OSName detects name of the operating system.
func OSName() string {
	f, err := os.Open("/etc/os-release")
	if err != nil {
		return "Linux (Unknown Distribution)"
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	name := "Linux (Unknown Distribution)"
	for s.Scan() {
		parts := strings.SplitN(s.Text(), "=", 2)
		switch parts[0] {
		case "Name":
			name = strings.Trim(parts[1], "\"")
		}
	}
	return name
}

// OSVersion detects version of the operating system.
func OSVersion() string {
	f, err := os.Open("/etc/os-release")
	if err != nil {
		return unknown
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	version := unknown
	for s.Scan() {
		parts := strings.SplitN(s.Text(), "=", 2)
		switch parts[0] {
		case "VERSION":
			version = strings.Trim(parts[1], "\"")
		case "VERSION_ID":
			if version == "" {
				version = strings.Trim(parts[1], "\"")
			}
		}
	}
	return version
}
