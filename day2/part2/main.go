package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"log/slog"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	if err != nil {
		logger.Error(`Failed to open file`, `error`, err)
		os.Exit(1)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	total := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		total += gameIdIfValidOtherwiseZero(logger, line)
	}

	fmt.Println(total)
}

func gameIdIfValidOtherwiseZero(logger *slog.Logger, line string) int {
	gameAndValues := strings.Split(line, `:`)

	_, err := strconv.Atoi(gameAndValues[0][5:])
	if err != nil {
		logger.Error(`Failed to convert str to number`, `string`, gameAndValues)
		os.Exit(1)
	}

	maxMapping := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	valueLines := strings.Split(gameAndValues[1], `;`)
	colors := []string{"red", "green", "blue"}
	for _, valueLine := range valueLines {
		mapping := valueLineMapping(logger, valueLine)
		for _, color := range colors {
			maxMapping[color] = int(math.Max(float64(maxMapping[color]), float64(mapping[color])))
		}
	}

	power := maxMapping["red"] * maxMapping["green"] * maxMapping["blue"]
	logger.Info(`current line`, `line`, line, `power`, power, `maxMapping`, maxMapping)

	return power
}

func valueLineMapping(logger *slog.Logger, valueLine string) map[string]int {
	values := strings.Split(valueLine, `,`)
	mapping := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, value := range values {
		re := regexp.MustCompile(`\s*(\d+)\s(red|green|blue)`)
		matches := re.FindStringSubmatch(value)
		k := matches[2]
		v := matches[1]
		i, err := strconv.Atoi(v)
		if err != nil {
			logger.Error(`Failed to parse value`, `valueStr`, matches[1])
			os.Exit(1)
		}

		mapping[k] = i
	}

	return mapping
}
