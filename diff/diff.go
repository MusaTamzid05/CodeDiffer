package diff

import (
	"fmt"
)

type Diff struct {
	FirstLineContainer *LineDataContainer
	SecondLineContainer *LineDataContainer
}

// This LineDiff struct is used
// to contain the difference in
// each line.


type LineDiff struct {
	missMatchString1 string
	missMatchString2 string
}


func NewLineDiff(line1, line2 string) *LineDiff {
	lineDiff := LineDiff{missMatchString1 : line1, missMatchString2 : line2}
	return &lineDiff

}



// this struct is used for showing final diff result.
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
