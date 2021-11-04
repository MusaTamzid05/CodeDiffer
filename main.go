package main

import (
	"code_diff/diff"
	//"log"
	"fmt"
	//"flag"
)


func main() {
	/*
	filePathPtr := flag.String("filePath", "./test.txt", "file path for reading code")
	flag.Parse()


	lineContrainer, err := diff.NewLineContainer(*filePathPtr)

	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println(lineContrainer.String())
	*/

	preprocessor := diff.Preprocessor{}
	fmt.Println(preprocessor.Process("let var_data = 12"))
	fmt.Println(preprocessor.Process("def test_data(val_data, val_2)"))
	fmt.Println(preprocessor.Process("__data"))
}
