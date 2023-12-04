package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var total int64 = 0
	digitsInLetters := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	last := false
	for !last {
		lineRaw, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				last = true
			} else {
				panic(err)
			}
		}
		line := string(lineRaw)
		var i int64
		var firstNumberIndex int = len(line)
		var lastNumberIndex int = -1
		var firstDigitInLettersIndex int = len(line)
		var selectedFirstDigitLettersValue int64 = -1
		var lastDigitInLettersIndex int = -1
		var selectedLastDigitLettersValue int64 = -1
		for i = 1; i < 10; i++ {
			curFirstIndex := strings.Index(line, strconv.FormatInt(i, 10))
			if curFirstIndex != -1 && curFirstIndex < firstNumberIndex {
				firstNumberIndex = curFirstIndex
			}
			curLastIndex := strings.LastIndex(line, strconv.FormatInt(i, 10))
			if curLastIndex != -1 && curLastIndex > lastNumberIndex {
				lastNumberIndex = curLastIndex
			}
			curFirstLettersIndex := strings.Index(line, digitsInLetters[i-1])
			if curFirstLettersIndex != -1 && curFirstLettersIndex < firstDigitInLettersIndex {
				firstDigitInLettersIndex = curFirstLettersIndex
				selectedFirstDigitLettersValue = i
			}
			curLastLettersIndex := strings.LastIndex(line, digitsInLetters[i-1])
			if curLastLettersIndex != -1 && curLastLettersIndex > lastDigitInLettersIndex {
				lastDigitInLettersIndex = curLastLettersIndex
				selectedLastDigitLettersValue = i
			}
		}

		firstDigit := ""
		lastDigit := ""
		if firstNumberIndex != len(line) {
			firstDigit = string(line[firstNumberIndex])
		}
		if firstDigitInLettersIndex != len(line) {
			if firstDigit == "" || firstDigitInLettersIndex < firstNumberIndex {
				firstDigit = strconv.FormatInt(selectedFirstDigitLettersValue, 10)
			}
		}
		if lastNumberIndex != -1 {
			lastDigit = string(line[lastNumberIndex])
		}
		if lastDigitInLettersIndex != -1 {
			if lastDigit == "" || lastDigitInLettersIndex > lastNumberIndex {
				lastDigit = strconv.FormatInt(selectedLastDigitLettersValue, 10)
			}
		}

		digits, err := strconv.ParseInt(firstDigit+lastDigit, 10, 64)
		if err != nil {
			panic(err)
		}
		total += digits
	}
	fmt.Println(total)
}
