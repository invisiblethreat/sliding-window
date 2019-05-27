package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func help() {
	fmt.Printf("Supply 3 arguments: file, starting regex, ending regex\n")
	fmt.Printf("'-i' is an optional first argument to include the trigger lines\n")
}

func main() {

	if len(os.Args) < 4 {
		help()
		os.Exit(1)
	}

	include := false
	var posArgs []string
	// Lazy 'shift'
	if strings.TrimSpace(os.Args[1]) == "-i" {

		include = true
		posArgs = os.Args[2:]

	} else {
		posArgs = os.Args[1:]

	}

	lines, err := slidingWindow(posArgs[0], posArgs[1], posArgs[2], include)
	if err != nil {
		fmt.Printf("Something went wrong: %s\n", err.Error())
		os.Exit(1)
	}
	for _, line := range *lines {
		fmt.Println(line)
	}
}

// slidingWindow does the heavy lifting. Returning a slice makes this testble
func slidingWindow(file, startTag, endTag string, include bool) (*[]string, error) {
	wantedLines := []string{}
	loadedFile, err := os.Open(file)
	if err != nil {
		return &wantedLines, err
	}
	defer loadedFile.Close()

	start, err := regexp.Compile(startTag)
	if err != nil {
		return &wantedLines, err
	}
	end, err := regexp.Compile(endTag)
	if err != nil {
		return &wantedLines, err
	}

	printing := true
	scan := bufio.NewScanner(loadedFile)
	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
		if start.Match([]byte(line)) {
			printing = false
			if include {
				wantedLines = append(wantedLines, line)
			}
		}
		if printing {
			wantedLines = append(wantedLines, line)
		}
		if end.Match([]byte(line)) {
			printing = true
			if include {
				wantedLines = append(wantedLines, line)
			}

		}
	}
	return &wantedLines, nil
}
