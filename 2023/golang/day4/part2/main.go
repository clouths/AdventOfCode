package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type stack[T any] struct {
	Push   func(T)
	Pop    func() T
	Length func() int
}

type Card struct {
	index        int
	winningCount int
}

func Stack[T any]() stack[T] {
	slice := []T{}
	return stack[T]{
		Push: func(i T) {
			slice = append(slice, i)
		},
		Pop: func() T {
			res := slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			return res
		},
		Length: func() int {
			return len(slice)
		},
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	end := false
	cardCount := 0
	index := 0
	originalCards := []*Card{}
	cards := Stack[*Card]()
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
				}
			}
		}
		card := &Card{
			index:        index,
			winningCount: winningCount,
		}
		originalCards = append(originalCards, card)
		cards.Push(card)
		index++
	}
	for cards.Length() > 0 {
		cardCount++
		card := cards.Pop()
		for i := 1; i <= card.winningCount; i++ {
			cards.Push(originalCards[card.index+i])
		}
	}
	fmt.Println(cardCount)
}
