package diff

import (
	"strings"
)

type Preprocessor struct {

}

func NewPreprocessor() *Preprocessor {
	return &Preprocessor{}
}

func (p* Preprocessor) Process(line string) string  {
	newLine := ""
	capitalizeChar := false


	for _, char := range  line {
		// convert 'test_data' to 'testData'
		// this is the reason we use capitalizeChar
		// for

		tempStr := string(char) // can be optimized

		if capitalizeChar {
			tempStr = strings.ToUpper(tempStr)
			capitalizeChar = false
		}

		if tempStr == "_" {
			capitalizeChar = true
			continue
		}
		newLine += tempStr
	}


	return newLine
}
