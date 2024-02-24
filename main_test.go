package main

import (
	"bytes"
	"testing"
)

const testWithFilesResult = `├───file1.txt (25b)
├───folder1
│	├───file2.txt (25b)
│	└───file3.txt (empty)
└───folder2
	├───file4.txt (20b)
	└───folder3
		├───folder4
		│	└───file6.txt (34b)
		└───folder5
			└───file5.txt (25b)
`

func TestTreeWithFiles(t *testing.T) {
	out := new(bytes.Buffer)
	err := dirTree(out, "testdata", true)
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != testWithFilesResult {
		t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, testWithFilesResult)
	}
}

const testWithoutFilesResult = `├───folder1
└───folder2
	└───folder3
		├───folder4
		└───folder5
`

func TestTreeWithoutFiles(t *testing.T) {
	out := new(bytes.Buffer)
	err := dirTree(out, "testdata", false)
	if err != nil {
		t.Errorf("test for OK Failed - error")
	}
	result := out.String()
	if result != testWithoutFilesResult {
		t.Errorf("test for OK Failed - results not match\nGot:\n%v\nExpected:\n%v", result, testWithoutFilesResult)
	}
}
