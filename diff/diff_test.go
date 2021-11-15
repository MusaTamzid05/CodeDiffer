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
