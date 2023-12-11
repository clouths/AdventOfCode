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
	gearX    int
	gearY    int
	checked  bool
}

func addIfNumberExist(line int, position int, buffer *[]byte, numbers *[]*NumberInfo) {
	if len(*buffer) > 0 {
		number, err := strconv.ParseInt(string(*buffer), 10, 64)
		if err != nil {
			panic(err)
		}
		*numbers = append(*numbers, &NumberInfo{
			number:   int(number),
			size:     len(*buffer),
			line:     line,
			position: position - len(*buffer),
			gearX:    -1,
			gearY:    -1,
		})
		*buffer = []byte{}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	end := false
	numbers := []*NumberInfo{}
	gears := map[int]map[int]bool{}
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
		gears[count] = map[int]bool{}
		curWidth := len(line)
		if curWidth > maxWidth {
			maxWidth = curWidth
		}
		buffer := []byte{}
		for position, char := range strings.TrimSpace(string(line)) {
			switch char {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				buffer = append(buffer, byte(char))
			case '*':
				gears[count][position] = true
				fallthrough
			default:
				addIfNumberExist(count, position, &buffer, &numbers)
			}
		}
		addIfNumberExist(count, len(line)-1, &buffer, &numbers)
		count++
	}
	sumGearRatio := 0
	for _, number := range numbers {
		// check above
		curLine := number.line - 1
		curPosition := number.position - 1
		if curLine >= 0 {
			// left diagonal
			if curPosition >= 0 && gears[curLine][curPosition] {
				number.gearX = curPosition
				number.gearY = curLine
				continue
			}
			// above
			curPosition = number.position
			isPart := false
			for i := 0; i < number.size; i++ {
				if gears[curLine][curPosition+i] {
					number.gearX = curPosition + i
					number.gearY = curLine
					isPart = true
				}
			}
			if isPart {
				continue
			}
			// right diagonal
			curPosition = number.position + number.size
			if curPosition < maxWidth && gears[curLine][curPosition] {
				number.gearX = curPosition
				number.gearY = curLine
				continue
			}
		}

		// check same line
		curLine = number.line
		// before
		curPosition = number.position - 1
		if curPosition > 0 && gears[curLine][curPosition] {
			number.gearX = curPosition
			number.gearY = curLine
			continue
		}
		// after
		curPosition = number.position + number.size
		if curPosition < maxWidth && gears[curLine][curPosition] {
			number.gearX = curPosition
			number.gearY = curLine
			continue
		}

		// check below
		curLine = number.line + 1
		curPosition = number.position - 1
		if curLine < count {
			// left diagonal
			if curPosition >= 0 && gears[curLine][curPosition] {
				number.gearX = curPosition
				number.gearY = curLine
				continue
			}
			// below
			curPosition = number.position
			isPart := false
			for i := 0; i < number.size; i++ {
				if gears[curLine][curPosition+i] {
					number.gearX = curPosition + i
					number.gearY = curLine
					isPart = true
				}
			}
			if isPart {
				continue
			}
			// right diagonal
			curPosition = number.position + number.size
			if curPosition < maxWidth && gears[curLine][curPosition] {
				number.gearX = curPosition
				number.gearY = curLine
				continue
			}
		}
	}
	for i, numberA := range numbers {
		for j, numberB := range numbers {
			if i == j {
				continue
			}
			if numberA.gearX != -1 && !numberA.checked && numberB.gearX != -1 && !numberB.checked {
				if numberA.gearX == numberB.gearX && numberA.gearY == numberB.gearY {
					sumGearRatio += numberA.number * numberB.number
					numberA.checked = true
					numberB.checked = true
				}
			}
		}
	}

	fmt.Println(sumGearRatio)
}
