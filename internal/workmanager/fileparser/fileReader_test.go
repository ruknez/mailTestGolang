package fileparser

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
)

func TestCalculate(t *testing.T) {
	magicWord := "Go"

	testDataSlice := []struct {
		fileData string
		count    int
	}{
		{
			fileData: `test Go date
						fsdf
							fsdf Go
							sdfsdf`,
			count: 2,
		},
		{
			fileData: "",
			count:    0,
		},
		{
			fileData: `test G0 date
				sdfsdf Go`,
			count: 1,
		},
	}

	for _, td := range testDataSlice {
		num, err := helperWrapper(t, td.fileData, magicWord)
		if err != nil {
			t.Errorf("testWrapper err is not nil %v", err)
		}
		if num != td.count {
			t.Errorf("testWrapper num %d != count %d", num, td.count)
		}
	}
}

func helperWrapper(t *testing.T, fileData, magicWord string) (int, error) {
	t.Helper()
	tFile, err := ioutil.TempFile("", "tmpSoursFile")
	if err != nil {
		t.Errorf("cannot creat tmpfile %v", err)
		return 0, err
	}
	defer os.Remove(tFile.Name())
	defer tFile.Close()

	tFile.WriteString(fileData)

	fParser := NewFileParser(context.Background())
	return fParser.Calculate(tFile.Name(), magicWord)
}
