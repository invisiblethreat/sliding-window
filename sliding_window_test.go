package main

import (
	"testing"
)

const (
	testFile    = "data/test.txt"
	include     = true
	dontInclude = false
	start       = "start"
	end         = "end"
	// only the first and last lines should be returned
	without = 2
	// this result includes the matching lines that toggle state
	with = 4
	// negative case that should not occur
	unexpected = 3
	// this will compile as a regexp
	goodRegexp = "^start$"
	// this will not compile as a regexp
	badRegexp = "["
)

func TestSlidingWindowResults(t *testing.T) {

	lines, err := slidingWindow(testFile, start, end, dontInclude)
	if err != nil {
		t.Errorf("Unexpected error in slidingWindow: %s", err.Error())
	}
	if len(lines) != without {
		t.Errorf("Expected %d lines, but got %d lines", without, len(lines))
	}

	lines, err = slidingWindow(testFile, start, end, include)
	if err != nil {
		t.Errorf("Unexpected error in slidingWindow: %s", err.Error())
	}
	if len(lines) != with {
		t.Errorf("Expected %d lines, but got %d lines", without, len(lines))
	}

	lines, err = slidingWindow(testFile, start, end, include)
	if err != nil {
		t.Errorf("Unexpected error in slidingWindow: %s", err.Error())
	}
	if len(lines) == unexpected {
		t.Errorf("Expected %d lines, but got %d lines", without, len(lines))
	}
}

func TestSlidingWindowRegex(t *testing.T) {

	_, err := slidingWindow(testFile, goodRegexp, end, dontInclude)
	if err != nil {
		t.Errorf("%s failed to compile as regexp: %s", goodRegexp, err.Error())
	}
	_, err = slidingWindow(testFile, badRegexp, end, dontInclude)
	if err == nil {
		t.Errorf("Bad regexp compiled: %s", badRegexp)
	}
}
