package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	args := loadFlags()

	config, err := readConfig(args.fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-2)
	}

	var regFilterName, regFilterIndex *regexp.Regexp
	if args.filterName != "" {
		regFilterName, err = regexp.Compile(args.filterName)
		if err != nil {
			fmt.Println(err)
			fmt.Println(ErrInvalidFilterName)
			os.Exit(-3)
		}
	}
	if args.filterIndex != "" {
		regFilterIndex, err = regexp.Compile(args.filterIndex)
		if err != nil {
			fmt.Println(err)
			fmt.Println(ErrInvalidFilterIndex)
			os.Exit(-4)
		}
	}

	for _, entry := range config {
		value, err := strconv.ParseFloat(entry.Value, 64)
		if err != nil {
			fmt.Println(err)
			continue
		}
		absDefault := entry.Default
		absValue := value
		diff := math.Abs(absDefault - absValue)

		if regFilterName != nil && regFilterName.Match([]byte(entry.Name)) {
			fmt.Println(entry.toString())
		} else if regFilterIndex != nil && regFilterIndex.Match([]byte(strconv.Itoa(entry.ParamIndex))) {
			fmt.Println(entry.toString())
		} else if diff > args.difference {
			fmt.Println(entry.toString())
		}
	}
}
