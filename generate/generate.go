package main

// see generate.go in the directory above. this needs to be in a separate
// directory so that it's `go run`-able by `go generate`.

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	mainTemplate = `package main

// Code generated - DO NOT EDIT.

import "github.com/{{username}}/oracles-randomizer/randomizer"

func main() {
	randomizer.Main()
}
`
	versionTemplate = `package randomizer

// Code generated - DO NOT EDIT.

const version = {{version}}
`
)

var (
	usernamePattern = strings.ReplaceAll(
		filepath.FromSlash(`github.com/(.+)/oracles-randomizer`), `\`, `\\`)
	usernameRegexp = regexp.MustCompile(usernamePattern)
	versionRegexp  = regexp.MustCompile(`/(.+)-\d+-g(.+)`)
)

func main() {
	generateMain()
	generateVersion()
}

func generateMain() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	matches := usernameRegexp.FindStringSubmatch(wd)
	if matches == nil {
		panic("error getting import path from working directory")
	}

	s := strings.ReplaceAll(mainTemplate, "{{username}}", matches[1])
	if err := ioutil.WriteFile("main.go", []byte(s), 0644); err != nil {
		panic(err)
	}
}

func generateVersion() {
	// use git state as basis for version string
	describeCmd := exec.Command("git", "describe", "--all")
	output, err := describeCmd.Output()
	if err != nil {
		panic(err)
	}

	// output should be "tags/[tag]" or "heads/[branch]"
	version := fmt.Sprintf(`"%s"`,
		strings.SplitN(strings.TrimSpace(string(output)), "/", 2)[1])

	s := strings.ReplaceAll(versionTemplate, "{{version}}", version)
	err = ioutil.WriteFile("randomizer/version.go", []byte(s), 0644)
	if err != nil {
		panic(err)
	}
}
