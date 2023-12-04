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
		for i = 1; i < 10; i++ {
			curFirstIndex := strings.Index(line, strconv.FormatInt(i, 10))
			if curFirstIndex != -1 && curFirstIndex < firstNumberIndex {
				firstNumberIndex = curFirstIndex
			}
			curLastIndex := strings.LastIndex(line, strconv.FormatInt(i, 10))
			if curLastIndex != -1 && curLastIndex > lastNumberIndex {
				lastNumberIndex = curLastIndex
			}
		}
		digits, err := strconv.ParseInt(string([]byte{line[firstNumberIndex], line[lastNumberIndex]}), 10, 64)
		if err != nil {
			panic(err)
		}
		total += digits
	}
	fmt.Println(total)
}
