package main

import "flag"

func loadFlags() flags {
	flags := flags{}
	flag.StringVar(&flags.fileName, "i", "", "path to file, usually has the extension DHP or JSON  (required)")
	flag.StringVar(&flags.filterName, "filterName", "", "Regular expression to filter for parameter names (optional)")
	flag.StringVar(&flags.filterIndex, "filterIndex", "", "Regular expression to filter for parameter indexes (optional)")
	flag.Float64Var(&flags.difference, "difference", 0.001, "Display parameters whose difference (Default and actual value) is greater than this")

	flag.Parse()

	return flags
}
