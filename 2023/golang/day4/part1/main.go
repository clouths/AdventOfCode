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
	end := false
	totalPoints := 0
	for !end {
		rawLine, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				end = true
			} else {
				panic(err)
			}
		}
		line := strings.TrimSpace(string(rawLine))
		winning := []int{}
		have := []int{}
		winningCount := 0
		lastWinningPoints := 0
		points := 0
		split := strings.Split(strings.Split(line, ":")[1], "|")
		rawWinnings := split[0]
		rawHaves := split[1]
		for _, rawWinnig := range strings.Split(rawWinnings, " ") {
			if rawWinnig == "" {
				continue
			}
			curWin, err := strconv.ParseInt(rawWinnig, 10, 64)
			if err != nil {
				panic(err)
			}
			winning = append(winning, int(curWin))
		}
		for _, rawHave := range strings.Split(rawHaves, " ") {
			if rawHave == "" {
				continue
			}
			curHave, err := strconv.ParseInt(rawHave, 10, 64)
			if err != nil {
				panic(err)
			}
			have = append(have, int(curHave))
		}
		for _, curHave := range have {
			for _, curWinning := range winning {
				if curHave == curWinning {
					winningCount++
					curPoints := 0
					switch winningCount {
					case 1, 2:
						lastWinningPoints = 1
					default:
						lastWinningPoints *= 2
					}
					curPoints = lastWinningPoints
					points += curPoints

				}
			}
		}
		totalPoints += points
	}
	fmt.Println(totalPoints)
}
