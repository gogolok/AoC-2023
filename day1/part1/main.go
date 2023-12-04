package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

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

		re := regexp.MustCompile(`\d`)
		matches := re.FindAllString(line, -1)

		firstDigitStr := matches[0]
		lastDigitStr := matches[len(matches)-1]

		twoDigitNumberStr := firstDigitStr + lastDigitStr

		i, err := strconv.Atoi(twoDigitNumberStr)
		if err != nil {
			logger.Error(`Failed to convert str to number`, `string`, twoDigitNumberStr)
			os.Exit(1)
		}

		total += i

		logger.Info(`current line`, `line`, line, `first digit`, firstDigitStr, `last digit`, lastDigitStr, `two-digit number`, twoDigitNumberStr, `total`, total)
	}

	fmt.Println(total)
}
