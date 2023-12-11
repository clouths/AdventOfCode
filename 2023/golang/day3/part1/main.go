package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type NumberInfo struct {
	number   int
	size     int
	line     int
	position int
}

func addIfNumberExist(line int, position int, buffer *[]byte, numbers *[]NumberInfo) {
	if len(*buffer) > 0 {
		number, err := strconv.ParseInt(string(*buffer), 10, 64)
		if err != nil {
			panic(err)
		}
		*numbers = append(*numbers, NumberInfo{
			number:   int(number),
			size:     len(*buffer),
			line:     line,
			position: position - len(*buffer),
		})
		*buffer = []byte{}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	end := false
	numbers := []NumberInfo{}
	symbols := map[int]map[int]bool{}
	count := 0
	maxWidth := 0
	for !end {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				end = true
			} else {
				panic(err)
			}
		}
		symbols[count] = map[int]bool{}
		curWidth := len(line)
		if curWidth > maxWidth {
			maxWidth = curWidth
		}
		buffer := []byte{}
		for position, char := range strings.TrimSpace(string(line)) {
			switch char {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				buffer = append(buffer, byte(char))
			case '.':
				addIfNumberExist(count, position, &buffer, &numbers)
			default:
				addIfNumberExist(count, position, &buffer, &numbers)
				symbols[count][position] = true
			}
		}
		addIfNumberExist(count, len(line)-1, &buffer, &numbers)
		count++
	}
	sumPartNumber := 0
	for _, number := range numbers {
		// check above
		curLine := number.line - 1
		curPosition := number.position - 1
		if curLine >= 0 {
			// left diagonal
			if curPosition >= 0 && symbols[curLine][curPosition] {
				sumPartNumber += number.number
				continue
			}
			// above
			curPosition = number.position
			isPart := false
			for i := 0; i < number.size; i++ {
				if symbols[curLine][curPosition+i] {
					sumPartNumber += number.number
					isPart = true
				}
			}
			if isPart {
				continue
			}
			// right diagonal
			curPosition = number.position + number.size
			if curPosition < maxWidth && symbols[curLine][curPosition] {
				sumPartNumber += number.number
				continue
			}
		}

		// check same line
		curLine = number.line
		// before
		curPosition = number.position - 1
		if curPosition > 0 && symbols[curLine][curPosition] {
			sumPartNumber += number.number
			continue
		}
		// after
		curPosition = number.position + number.size
		if curPosition < maxWidth && symbols[curLine][curPosition] {
			sumPartNumber += number.number
			continue
		}

		// check below
		curLine = number.line + 1
		curPosition = number.position - 1
		if curLine < count {
			// left diagonal
			if curPosition >= 0 && symbols[curLine][curPosition] {
				sumPartNumber += number.number
				continue
			}
			// below
			curPosition = number.position
			isPart := false
			for i := 0; i < number.size; i++ {
				if symbols[curLine][curPosition+i] {
					sumPartNumber += number.number
					isPart = true
				}
			}
			if isPart {
				continue
			}
			// right diagonal
			curPosition = number.position + number.size
			if curPosition < maxWidth && symbols[curLine][curPosition] {
				sumPartNumber += number.number
				continue
			}
		}
	}
	fmt.Println(sumPartNumber)
}
