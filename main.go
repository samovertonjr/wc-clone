package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	byteCountPointer := flag.String("c", "", "a file")
	numberOfLinesPointer := flag.String("l", "", "a file")
	flag.Parse()

	if *byteCountPointer != "" {
		getFileSizeInBytes(byteCountPointer)
		os.Exit(0)
	}

	if *numberOfLinesPointer != "" {
		getNumberOfLines(numberOfLinesPointer)
		os.Exit(0)
	}

}

func getFileSizeInBytes(file *string) {
	fileInfo, err := os.Stat(*file)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fileInfo, ":", *file)
}

func getNumberOfLines(file *string) {
	fileInfo, err := os.Open(*file)

	if err != nil {
		log.Fatal(err)
	}

	defer fileInfo.Close()

	scanner := bufio.NewScanner(fileInfo)

	lineNumber := 1

	for scanner.Scan() {
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(lineNumber, ":", *file)
}
