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

	stats := make([]int, 4)

	for _, option := range options {
		switch option {
		case "-c", "--bytes":
			bytesCount, error := counters.CountBytes(path)

			if error != nil {
				log.Fatal(error)
			}

			stats[0] = bytesCount
		case "-m", "--chars":
			charsCount := counters.CountChars(path)
			stats[1] = charsCount
		case "-l", "--lines":
			linesCount := counters.CountLines(path)
			stats[2] = linesCount
		case "-w", "--words":
			wordsCount := counters.CountWords(path)
			stats[3] = wordsCount
		default:
			log.Fatal("Incorrect option")
		}
	}

	fmt.Printf("%d %d %d %s \n", stats[0], stats[1], stats[2], path)
}
