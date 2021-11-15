package diff

import (
	"fmt"
	"strings"
	"regexp"
)

type LineData struct {
	Line string
	OriginalIndex int // index before removing space
	NewIndex int // index after removing index
}

func (l *LineData) String() string {
	str := fmt.Sprintf("Line : %s\n", l.Line)
	str += fmt.Sprintf("Original Index : %d\n", l.OriginalIndex)
	str += fmt.Sprintf("New Index : %d\n", l.NewIndex)

	return str
}



type LineDataContainer struct {
	LineData [] *LineData
}


func NewLineData(Line string, OriginalIndex , NewIndex int) *LineData {
	return &LineData{Line: Line, NewIndex : NewIndex, OriginalIndex : OriginalIndex}
}

func NewLineContainer(filePath string) (*LineDataContainer, error){
	lines, err  := Read(filePath)

	if err != nil {
		return nil, err
	}

	lineData := []*LineData{}
	newIndex := 0

	for index, line := range lines {
		trimLine := strings.Trim(line, " ")
		
		if len(trimLine) == 0 {
			continue
		}

		lineData = append(lineData, NewLineData(line, index, newIndex))
		newIndex += 1
	}

	lineDataContainer := LineDataContainer{}
	lineDataContainer.LineData = lineData

	return &lineDataContainer, nil

}


func (l *LineDataContainer) String() string {
	str := ""

	for _, lineData := range l.LineData {
		str += lineData.String()
	}

	return str
}

func (l *LineDataContainer) isFilePath(str string) bool {
	match, _ := regexp.MatchString("\\.(\\w+)", str)
	return match
}
