package main

import (
	"bufio"
	"fmt"
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

		re := regexp.MustCompile("\\d|one|two|three|four|five|six|seven|eight|nine")
		matches := re.FindAllString(line, -1)
		firstDigitStr := spelledDigitToDigit(matches[0])

		re = regexp.MustCompile(".*(\\d|one|two|three|four|five|six|seven|eight|nine)")
		matches = re.FindStringSubmatch(line)
		lastDigitStr := spelledDigitToDigit(matches[1])

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

func spelledDigitToDigit(s string) string {
	s = strings.Replace(s, `one`, `1`, -1)
	s = strings.Replace(s, `two`, `2`, -1)
	s = strings.Replace(s, `three`, `3`, -1)
	s = strings.Replace(s, `four`, `4`, -1)
	s = strings.Replace(s, `five`, `5`, -1)
	s = strings.Replace(s, `six`, `6`, -1)
	s = strings.Replace(s, `seven`, `7`, -1)
	s = strings.Replace(s, `eight`, `8`, -1)
	s = strings.Replace(s, `nine`, `9`, -1)
	return s
}
