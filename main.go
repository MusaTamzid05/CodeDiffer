package main

import (
	"code_diff/diff"
	"log"
	"fmt"
)


func main() {
	lineContrainer, err := diff.NewLineContainer("./test.txt")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(lineContrainer.String())

}
