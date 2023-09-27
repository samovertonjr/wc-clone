package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	byteCountPointer := flag.String("c", "", "a file")
	numberOfLinesPointer := flag.String("l", "", "a file")
	numberOfWordsPointer := flag.String("w", "", "a file")
	numberOfCharsPointer := flag.String("m", "", "a file")

	flag.Parse()

	if *byteCountPointer != "" {
		getFileSizeInBytes(byteCountPointer)
		os.Exit(0)
	}

	if *numberOfLinesPointer != "" {
		getNumberOfLines(numberOfLinesPointer)
		os.Exit(0)
	}

	if *numberOfWordsPointer != "" {
		getNumberOfWords(numberOfWordsPointer)
		os.Exit(0)
	}

	if *numberOfCharsPointer != "" {
		getNumberOfChars(numberOfCharsPointer)
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

func getNumberOfWords(file *string) {
	fileInfo, err := os.Open(*file)

	if err != nil {
		log.Fatal(err)
	}

	defer fileInfo.Close()

	scanner := bufio.NewScanner(fileInfo)

	wordCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		wordCount = wordCount + len(words)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(wordCount, ":", *file)
}

func getNumberOfChars(file *string) {
	fileInfo, err := os.Open(*file)

	if err != nil {
		log.Fatal(err)
	}

	defer fileInfo.Close()

	scanner := bufio.NewScanner(fileInfo)

	characterCount := 0

	for scanner.Scan() {

		line := scanner.Text()
		characterCount += len(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(characterCount, ":", *file)
}
