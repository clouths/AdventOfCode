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

type SeedInfo struct {
	rangeStart  int
	rangeLength int
}

type Puzzle struct {
	seedInfos   []SeedInfo
	convertMaps []*[]ConvertMap
}

func getSeeds(input string) (*[]SeedInfo, error) {
	seeds := &[]SeedInfo{}
	rawSeeds := strings.Split(input, ":")[1]
	partNumber := 0
	var seedInfo *SeedInfo
	for _, rawSeed := range strings.Split(rawSeeds, " ") {
		if rawSeed == "" {
			continue
		}
		seedPart, err := strconv.ParseInt(rawSeed, 10, 64)
		if err != nil {
			return nil, err
		}
		if partNumber == 0 {
			partNumber = 1
			seedInfo = &SeedInfo{}
			seedInfo.rangeStart = int(seedPart)
		} else {
			partNumber = 0
			seedInfo.rangeLength = int(seedPart)
			*seeds = append(*seeds, *seedInfo)
		}
	}
	return seeds, nil
}

func readPuzzle() (*Puzzle, error) {
	var seeds []SeedInfo
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
		seedInfos:   seeds,
		convertMaps: convertMaps,
	}, nil
}

func seedsToSoils(puzzle *Puzzle) []int {
	soils := []int{}
	// checking one seedInfo at at time since it takes a lot of ram
	// improvement: run in parallel
	// improvement2: don't load the whole range before iterating to save ram
	for _, seedInfo := range puzzle.seedInfos {
		curSoils := []int{}
		rangeEnd := seedInfo.rangeStart + seedInfo.rangeLength
		for value := seedInfo.rangeStart; value < rangeEnd; value++ {
			curSoils = append(curSoils, value)
		}

		lowest := math.MaxInt
		for _, curMaps := range puzzle.convertMaps {
			for key, value := range curSoils {
				for _, curMap := range *curMaps {
					srcRangeEnd := curMap.srcRangeStart + curMap.rangeLength
					if value >= curMap.srcRangeStart && value < srcRangeEnd {
						diff := value - curMap.srcRangeStart
						curSoils[key] = curMap.dstRangeStart + diff
						break
					}
				}
			}
		}
		for _, soil := range curSoils {
			if soil < lowest {
				lowest = soil
			}
		}
		soils = append(soils, lowest)
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
