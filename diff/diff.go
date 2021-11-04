package diff

import (
	"fmt"
)

type Diff struct {
	FirstLineContainer *LineDataContainer
	SecondLineContainer *LineDataContainer
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

func (d *Diff) Run() {

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
		}


	}

}
