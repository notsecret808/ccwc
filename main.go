package main

import (
	"fmt"
	"log"
	"notsecret808/ccwc/counters"
	"notsecret808/ccwc/utils"
)

func main() {
	options, path, err := utils.ParseCmdParams()

	if err != nil {
		log.Fatal(err)
	}

	f := File{Path: path}

	if len(options) == 0 {
		f.countBytes()
		f.countChars()
		f.countLines()
		f.countWords()
	}

	for _, option := range options {
		switch option {
		case "-c", "--bytes":
			f.countBytes()
		case "-m", "--chars":
			f.countChars()
		case "-l", "--lines":
			f.countLines()
		case "-w", "--words":
			f.countWords()
		default:
			log.Fatal("Incorrect option")
		}
	}

	f.writeOutput()
}

type File struct {
	Path  string
	Bytes *int
	Chars *int
	Lines *int
	Words *int
}

func (f *File) countBytes() {
	bytesCount, error := counters.CountBytes(f.Path)

	if error != nil {
		log.Fatal(error)
	}

	if f.Bytes == nil {
		f.Bytes = new(int)
	}

	*f.Bytes = bytesCount
}

func (f *File) countChars() {
	charsCount, error := counters.CountChars(f.Path)

	if error != nil {
		log.Fatal(error)
	}

	if f.Chars == nil {
		f.Chars = new(int)
	}

	*f.Chars = charsCount

}

func (f *File) countLines() {
	linesCount, error := counters.CountLines(f.Path)

	if error != nil {
		log.Fatal(error)
	}

	if f.Lines == nil {
		f.Lines = new(int)
	}

	*f.Lines = linesCount
}

func (f *File) countWords() {
	wordsCount, error := counters.CountWords(f.Path)

	if error != nil {
		log.Fatal(error)
	}

	if f.Words == nil {
		f.Words = new(int)
	}

	*f.Words = wordsCount
}

func (f *File) writeOutput() {
	output := ""

	var values []*int

	values = append(values, f.Bytes)
	values = append(values, f.Chars)
	values = append(values, f.Lines)
	values = append(values, f.Words)

	for _, value := range values {
		if value != nil {
			output = fmt.Sprintf("%s %d", output, *value)
		}
	}

	fmt.Printf("%s %s\n", output, f.Path)
}
