package main

import (
	"code_diff/diff"
	"flag"
	"fmt"
)


func main() {
	firstFilePathPtr := flag.String("firstFile", "./sample/test5.py", "path for the first file")
	secondPathPtr := flag.String("secondFile", "./sample/test6.py", "file path for second file")
	flag.Parse()

	diff, err := diff.NewDiff(*firstFilePathPtr, *secondPathPtr)

	if err != nil {
		fmt.Println(err)
		return
	}

	result := diff.Run()

	fmt.Println(result.Count)

	for _, diff := range result.LineDiffs {
		fmt.Println(diff.Line1)
		fmt.Println(diff.Line2)
		fmt.Println()
	}

}
