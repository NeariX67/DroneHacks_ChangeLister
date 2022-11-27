package main

import "flag"

func loadFlags() flags {
	flags := flags{}
	flag.StringVar(&flags.fileName, "i", "", "path to file")
	flag.StringVar(&flags.filterName, "filterName", "", "")
	flag.StringVar(&flags.filterIndex, "filterIndex", "", "")
	flag.Float64Var(&flags.difference, "difference", 0.001, "")

	flag.Parse()

	return flags
}
