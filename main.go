package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	filePointer := flag.String("c", "test.txt", "a file")
	fileSize := getFileSize(filePointer)

	flag.Parse()

	fmt.Println(fileSize, ":", *filePointer)

	// fmt.Println("tail:", flag.Args())
}

func getFileSize(fileName *string) int64 {
	fileInfo, err := os.Stat(*fileName)

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return fileInfo.Size()
}
