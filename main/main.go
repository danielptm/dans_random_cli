package main

import (
	"dpt_cli/modules"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 6 {
		printHelp()
	} else {
		prog := args[1]
		firstFlag := args[2]
		firstValue := args[3]
		secondFlag := args[4]
		secondValue := args[5]
		if prog == "line-search" && firstFlag == "-s" && len(firstValue) > 0 && secondFlag == "-p" && len(secondValue) > 0 {
			res, e := modules.GetAllFilesForLocation(secondValue)
			if e == nil {
				for _, e := range res {
					res, e := modules.GetLinesThatMatchFromFile(e, firstValue)
					if e == nil {
						for _, e := range res {
							print(e)
						}
					}
				}
			}
		} else {
			printHelp()
		}
	}
}

func printHelp() {
	println("")
	println("Dan's random cli")
	println("Options:")
	println("line-search -s <search-string> -p <absolute/path/to/folder/location>")
	println("")
}
