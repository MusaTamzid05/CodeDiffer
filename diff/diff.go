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
	Line1 string
	Line2 string
}


func NewLineDiff(line1, line2 string) *LineDiff {
	lineDiff := LineDiff{Line1: line1, Line2 : line2}
	return &lineDiff
}

func (l *LineDiff) Show() {

	green := "\033[38;5;118m%s\033[39;49m\n"
	red := "\033[38;5;124m%s\033[39;49m\n"

	fmt.Printf(green,l.Line1)
	fmt.Printf(red,l.Line2)
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
			diffResult.Count += 1
			diffResult.LineDiffs = append(diffResult.LineDiffs, NewLineDiff( "> "+  firstLineData.String(), "< " + secondLineData.String()) )
		}
	}

	return diffResult

}
