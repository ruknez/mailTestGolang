package workmanager

import (
	"testing"
)

func TestIsItURL(t *testing.T) {
	testDataSlice := []struct {
		testData string
		isURL    bool
	}{
		{testData: "", isURL: false},
		{testData: "https://golang.org", isURL: true},
		{testData: "http://golang.org", isURL: true},
		{testData: "https://golang", isURL: false},
		{testData: "www.golang.org", isURL: false},
		{testData: "  https://golang.org ", isURL: false},
		{testData: "", isURL: false},
	}
	for _, td := range testDataSlice {
		td := td
		t.Run(td.testData, func(t *testing.T) {
			if isItURL(td.testData) != td.isURL {
				t.Errorf("error in %s res %v expect %v", td.testData, isItURL(td.testData), td.isURL)
			}
		})
	}
}
