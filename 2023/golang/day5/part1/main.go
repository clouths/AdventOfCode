package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Step int

const (
	Seed int = 0
	Map      = 1
)

type ConvertMap struct {
	dstRangeStart int
	srcRangeStart int
	rangeLength   int
}

type Puzzle struct {
	seeds       []int
	convertMaps []*[]ConvertMap
}

func getSeeds(input string) (*[]int, error) {
	seeds := &[]int{}
	rawSeeds := strings.Split(input, ":")[1]
	for _, rawSeed := range strings.Split(rawSeeds, " ") {
		if rawSeed == "" {
			continue
		}
		seed, err := strconv.ParseInt(rawSeed, 10, 64)
		if err != nil {
			return nil, err
		}
		*seeds = append(*seeds, int(seed))
	}
	return seeds, nil
}

func readPuzzle() (*Puzzle, error) {
	var seeds []int
	convertMaps := []*[]ConvertMap{}
	reader := bufio.NewReader(os.Stdin)
	var curConvertMap *[]ConvertMap
	step := Seed

	end := false
	for !end {
		rawLine, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				end = true
			} else {
				return nil, err
			}
		}
		line := strings.TrimSpace(string(rawLine))

		switch step {
		case Seed:
			pSeeds, err := getSeeds(line)
			if err != nil {
				return nil, err
			}
			seeds = *pSeeds
			step = Map
		case Map:
			if curConvertMap == nil {
				curConvertMap = &[]ConvertMap{}
			}
			if line == "" || strings.Contains(line, "map") {
				if len(*curConvertMap) > 0 {
					convertMaps = append(convertMaps, curConvertMap)
					curConvertMap = nil
				}
				continue
			}
			curMap := ConvertMap{}
			count, err := fmt.Sscan(line, &curMap.dstRangeStart, &curMap.srcRangeStart, &curMap.rangeLength)
			if err != nil {
				return nil, err
			}
			if count != 3 {
				return nil, errors.New("receive invalid map entry")
			}
			*curConvertMap = append(*curConvertMap, curMap)
		}
		if end {
			convertMaps = append(convertMaps, curConvertMap)
			curConvertMap = nil
		}
	}

	return &Puzzle{
		seeds:       seeds,
		convertMaps: convertMaps,
	}, nil
}

func seedsToSoils(puzzle *Puzzle) []int {
	soils := []int{}
	soils = append(soils, puzzle.seeds...)
	for _, curMaps := range puzzle.convertMaps {
		for key, value := range soils {
			for _, curMap := range *curMaps {
				srcRangeEnd := curMap.srcRangeStart + curMap.rangeLength
				if value >= curMap.srcRangeStart && value < srcRangeEnd {
					diff := value - curMap.srcRangeStart
					soils[key] = curMap.dstRangeStart + diff
					break
				}
			}
		}
	}
	return soils
}

func main() {
	puzzle, err := readPuzzle()
	if err != nil {
		panic(err)
	}

	soils := seedsToSoils(puzzle)

	lowest := math.MaxInt
	for _, soil := range soils {
		if soil < lowest {
			lowest = soil
		}
	}

	fmt.Println(lowest)
}
