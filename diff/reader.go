package diff

import (
	"os"
	"bufio"
)

func Read(filePath string)([]string, error){
	lines := []string{}

	fp, err := os.Open(filePath)

	if err != nil {
		return lines, err
	}

	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil

}

