package diff

import (
	"testing"
)


func TestDiffResultCount(t *testing.T) {

	diff, err := NewDiff("../sample/test.c", "../sample/test2.c")

	if err != nil {
		t.Errorf("Error opening file, Error Message :%s", err)
	}

	result := diff.Run()

	if result.Count != 1 {
		t.Errorf("Count found %d,expected %d", 1, result.Count)
	}

}


func TestDoesNotShowCaseDiffWhenTwoCaseDiff(t *testing.T) {

	diff, err := NewDiff("../sample/test3.c", "../sample/test4.c")

	if err != nil {
		t.Errorf("Error opening file, Error Message :%s", err)
	}

	result := diff.Run()

	if result.Count != 0 {
		t.Errorf("Count found %d,expected %d", 0, result.Count)
	}

}


func TestIsFilePathStr(t *testing.T) {
	diff := Diff{}

	if diff.isFilePath("normal string") {
		t.Errorf("String with no filepath detected as filepath")
	}

	if !diff.isFilePath("test.txt") {
		t.Errorf("filepath is not detected properly.")
	}
}
