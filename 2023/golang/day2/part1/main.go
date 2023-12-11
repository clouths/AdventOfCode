package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Games struct {
	id       int
	games    []Game
	maxRed   int
	maxGreen int
	maxBlue  int
}

type Game struct {
	red   int
	green int
	blue  int
}

func (g Games) isPossible(limit Game) bool {
	for _, game := range g.games {
		if game.red > g.maxRed {
			g.maxRed = game.red
		}
		if game.green > g.maxGreen {
			g.maxGreen = game.green
		}
		if game.blue > g.maxBlue {
			g.maxBlue = game.blue
		}
	}
	result := g.maxRed > limit.red || g.maxGreen > limit.green || g.maxBlue > limit.blue
	return !result
}

func newGames(input string) (*Games, error) {
	splitInput := strings.Split(input, ":")
	rawGameId := strings.Split(splitInput[0], " ")[1]
	gameId, err := strconv.ParseInt(rawGameId, 10, 64)
	if err != nil {
		return nil, err
	}
	rawGames := strings.Split(splitInput[1], ";")
	games := []Game{}
	for _, rawGame := range rawGames {
		rawColours := strings.Split(rawGame, ",")
		game := Game{}
		for _, rawColour := range rawColours {
			rawColourParts := strings.Split(rawColour, " ")
			count, err := strconv.ParseInt(rawColourParts[1], 10, 64)
			if err != nil {
				return nil, err
			}
			switch strings.TrimSpace(rawColourParts[2]) {
			case "red":
				game.red = int(count)
			case "green":
				game.green = int(count)
			case "blue":
				game.blue = int(count)
			}
		}
		games = append(games, game)
	}
	return &Games{
		id:    int(gameId),
		games: games,
	}, nil
}

func main() {
	end := false
	reader := bufio.NewReader(os.Stdin)
	total := 0
	resultGame := Game{
		red:   12,
		green: 13,
		blue:  14,
	}
	for !end {
		input, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				end = true
			} else {
				panic(err)
			}
		}
		games, err := newGames(string(input))
		if err != nil {
			panic(err)
		}
		if games.isPossible(resultGame) {
			total += games.id
		}
	}
	fmt.Println(total)
}
