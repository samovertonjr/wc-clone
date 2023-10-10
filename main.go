package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
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

	var data []byte
	var err error
	numberOfFlags := flag.NFlag()
	arg := flag.Arg(0)

	stat, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		data, err = io.ReadAll(os.Stdin)

	} else {
		data, err = os.ReadFile(arg)
	}

	if numberOfFlags == 0 {
		if err != nil {
			log.Fatal(err)
		}

		fileSizeInBytes := getFileSizeInBytes(data)
		numberOfLines := getNumberOfLines(data)
		numberOfWords := getNumberOfWords(data)
		numberOfChars := getNumberOfChars(data)

		fmt.Printf("%d Bytes \n", fileSizeInBytes)
		fmt.Printf("%d Lines \n", numberOfLines)
		fmt.Printf("%d Words \n", numberOfWords)
		fmt.Printf("%d Chars \n", numberOfChars)
		fmt.Printf("%s \n", flag.Arg(0))
		os.Exit(0)
	}

	if *byteCountPointer {
		fileSizeInBytes := getFileSizeInBytes(data)
		fmt.Printf("%d Bytes \n", fileSizeInBytes)
	}

	if *numberOfLinesPointer {
		numberOfLines := getNumberOfLines(data)
		fmt.Printf("%d Lines \n", numberOfLines)
	}

	if *numberOfWordsPointer {
		numberOfWords := getNumberOfWords(data)
		fmt.Printf("%d Words \n", numberOfWords)
	}

	if *numberOfCharsPointer {
		numberOfChars := getNumberOfChars(data)
		fmt.Printf("%d Chars \n", numberOfChars)
	}

	if flag.NArg() > 0 {
		fmt.Printf("%s \n", flag.Arg(0))
	}

	os.Exit(0)
}

func getFileSizeInBytes(data []byte) int {
	var size int
	reader := bytes.NewReader(data)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		size += len(scanner.Text()) + 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return size
}

func getNumberOfLines(data []byte) int {
	reader := bytes.NewReader(data)

	scanner := bufio.NewScanner(reader)

	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lineNumber
}

func getNumberOfWords(data []byte) int {
	reader := bytes.NewReader(data)

	scanner := bufio.NewScanner(reader)

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

func getNumberOfChars(data []byte) int {
	reader := bytes.NewReader(data)

	scanner := bufio.NewScanner(reader)

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
