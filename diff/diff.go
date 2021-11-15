package diff

import (
	"fmt"
	"regexp"
)

type Diff struct {
	FirstLineContainer *LineDataContainer
	SecondLineContainer *LineDataContainer
}

// This LineDiff struct is used
// to contain the difference in
// each line.

type LineDiffData struct {
	missMatchString1 string
	missMatchString2 string
}

type LineDiff struct {
	DiffData []*LineDiffData
}

func getMissMatch(index int, str1, str2 string) (int, *LineDiffData ){
	charIndex := index
	missMatchString1 := ""
	missMatchString2 := ""

	for string(str1[charIndex]) != string(str2[charIndex]) {
		missMatchString1 += string(str1[charIndex])
		missMatchString2 += string(str2[charIndex])
		charIndex += 1
	}

	return charIndex, &LineDiffData{missMatchString1 : missMatchString1, missMatchString2 : missMatchString2}
}

func NewLineDiff(line1, line2 string) *LineDiff {
	diffs := []*LineDiffData{}
	charIndex := 0
	missMatchLineData := &LineDiffData{}

	for charIndex < len(line1) {
		if line1[charIndex] != line2[charIndex] {
			charIndex, missMatchLineData = getMissMatch(charIndex, line1, line2)
			diffs = append(diffs, missMatchLineData)
			continue
		}
		charIndex += 1

	}



	lineDiff := LineDiff{DiffData : diffs}

	return &lineDiff
}


type DiffResult struct {
	Count int
	LineDiffs []*LineDiff
}

func NewDiffResult() *DiffResult {
	return &DiffResult{Count: 0}
}

func NewDiff(firstFilePath, secondFilePath string) (*Diff, error){
	firstLineContainer, err := NewLineContainer(firstFilePath)

	if err != nil {
		return nil, err
	}

	secondLineContainer , err := NewLineContainer(secondFilePath)

	if err != nil {
		return nil, err

	}

	return &Diff{FirstLineContainer : firstLineContainer , SecondLineContainer : secondLineContainer}, nil
}

func (d *Diff) isFilePath(str string) bool {
	match, _ := regexp.MatchString("\\.(\\w+)", str)
	return match
}

func (d *Diff) Run() *DiffResult  {

	diffResult := NewDiffResult()
	preprocessor := NewPreprocessor()

	for i, firstLineData := range d.FirstLineContainer.LineData {
		secondLineData := d.SecondLineContainer.LineData[i]

		firstProcessedLine := preprocessor.Process(firstLineData.Line)
		secondProcessedLine := preprocessor.Process(secondLineData.Line)

		if firstProcessedLine != secondProcessedLine {
			fmt.Println("========================")
			fmt.Println(firstLineData.String())
			fmt.Println(secondLineData.String())
			fmt.Println("========================")

			diffResult.Count += 1
			diffResult.LineDiffs = append(diffResult.LineDiffs, NewLineDiff(firstProcessedLine, secondProcessedLine) )
		}


	}

	return diffResult

}
