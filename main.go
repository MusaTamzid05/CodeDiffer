package main

import (
	"code_diff/diff"
	"log"
)


func main() {
	lines, err := diff.Read("./main.go")

	if err != nil {
		log.Fatalln(err)
	}

	for _, line := range lines {
		log.Println(line)
	}
}
