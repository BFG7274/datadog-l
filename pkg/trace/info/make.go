// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build ignore
// +build ignore

// The 'make' program is run by go generate to compile the versioning information
// into the info package. It expects the 'git' command to be installed.
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

// runs runs the given command and returns the output. If it fails,
// or if the result is an empty string, it returns the fallback.
func run(fallback, name string, args ...string) string {
	cmd := exec.Command(name, args...)
	out, err := cmd.Output()
	if err != nil || len(out) == 0 {
		return fallback
	}
	return string(bytes.Trim(out, "\n"))
}

func main() {
	log.SetPrefix("make_version")
	log.SetFlags(0)

	commit := run("master", "git", "rev-parse", "--short", "HEAD")
	branch := run("master", "git", "rev-parse", "--abbrev-ref", "HEAD")
	version := os.Getenv("TRACE_AGENT_VERSION")
	if version == "" {
		version = "0.99.0"
	}

	output := fmt.Sprintf(template, version, commit, branch, time.Now().String())
	err := ioutil.WriteFile("git_version.go", []byte(output), 0664)
	if err != nil {
		log.Fatal(err)
	}
}

const template = `// Code generated by 'go run make.go'. DO NOT EDIT.

package info

import (
	"runtime"
	"strings"
)

func init() {
	Version = %[1]q
	GitCommit = %[2]q
	GitBranch = %[3]q
	BuildDate = %[4]q
	GoVersion = strings.TrimPrefix(runtime.Version(), "go")
}
`
