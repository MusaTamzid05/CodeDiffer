package main

import (
	"code_diff/diff"
	"flag"
	"fmt"
)


func main() {
	firstFilePathPtr := flag.String("firstFile", "./sample/test.c", "path for the first file")
	secondPathPtr := flag.String("secondFile", "./sample/test2.c", "file path for second file")
	flag.Parse()

	diff, err := diff.NewDiff(*firstFilePathPtr, *secondPathPtr)

	if err != nil {
		fmt.Println(err)
		return
	}

	result := diff.Run()

	fmt.Println(result.Count)

}
