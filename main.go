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

	byteCountPointer := flag.Bool("c", false, "a file")
	numberOfLinesPointer := flag.Bool("l", false, "a file")
	numberOfWordsPointer := flag.Bool("w", false, "a file")
	numberOfCharsPointer := flag.Bool("m", false, "a file")

	flag.Parse()

	numberOfFlags := flag.NFlag()
	arg := flag.Arg(0)

	if numberOfFlags == 0 {
		fileSizeInBytes := getFileSizeInBytes(&arg)
		numberOfLines := getNumberOfLines(&arg)
		numberOfWords := getNumberOfWords(&arg)

		fmt.Printf("%o \n", fileSizeInBytes)
		fmt.Printf("%o \n", numberOfLines)
		fmt.Printf("%o \n", numberOfWords)
		fmt.Printf("%s \n", flag.Arg(0))
		os.Exit(0)
	}

	if *byteCountPointer {
		fileSizeInBytes := getFileSizeInBytes(&arg)
		fmt.Printf("%o \n", fileSizeInBytes)
	}

	if *numberOfLinesPointer {
		numberOfLines := getNumberOfLines(&arg)
		fmt.Printf("%o \n", numberOfLines)
	}

	if *numberOfWordsPointer {
		numberOfWords := getNumberOfWords(&arg)
		fmt.Printf("%o \n", numberOfWords)
	}

	if *numberOfCharsPointer {
		numberOfChars := getNumberOfChars(&arg)
		fmt.Printf("%o \n", numberOfChars)
	}

	if flag.NArg() > 0 {
		fmt.Printf("%s \n", flag.Arg(0))
	}

	os.Exit(0)
}

func getFileSizeInBytes(file *string) int64 {
	fileInfo, err := os.Stat(*file)

	if err != nil {
		log.Fatal(err)
	}

	return fileInfo.Size()
}

func getNumberOfLines(file *string) int {
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

	return lineNumber
}

func getNumberOfWords(file *string) int {
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

	return wordCount
}

func getNumberOfChars(file *string) int {
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

	return characterCount
}
