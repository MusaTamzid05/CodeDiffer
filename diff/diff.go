package diff

import (
	"fmt"
)

type Diff struct {
	FirstLineContainer *LineDataContainer
	SecondLineContainer *LineDataContainer
}

type DiffData struct {
	LineData1 *LineData
	LineData2 *LineData
}

func NewDiffData(lineData1, lineData2  *LineData) *DiffData {
	return &DiffData{LineData1 : lineData1, LineData2 : lineData2}
}

type DiffResult struct {
	Count int
	Data []*DiffData
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
			diffResult.Data = append(diffResult.Data, NewDiffData(firstLineData, secondLineData))
		}


	}

	return diffResult

}
