package main

import (
	"dpt_cli/modules"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		prog := args[1]
		if prog == "pretty-json" {
			if len(args) > 3 {
				firstFlag := args[2]
				firstValue := args[3]
				if firstFlag == "-s" && len(firstValue) > 0 {
					s := modules.GetPrettyJson(firstValue)
					modules.WriteFileWithContents(s, firstValue)
				}
			}
		}

		if prog == "line-search" {
			if len(args) > 5 {
				firstFlag := args[2]
				firstValue := args[3]
				secondFlag := args[4]
				secondValue := args[5]
				if firstFlag == "-p" && len(firstValue) > 0 && secondFlag == "-c" && len(secondValue) > 0 {
					res, e := modules.GetAllFilesForLocation(firstValue)
					println(fmt.Sprintf("There are %d instances of this string", len(res)))
					if e == nil {
						for _, e := range res {
							res, e := modules.GetLinesThatMatchFromFile(e, secondValue)
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
	} else {
		printHelp()
	}
}

func printHelp() {
	println("")
	println("Dan's random cli")
	println("")
	println("Options:")
	println("")
	println("line-search -p <absolute/path/to/folder/location> -c <search-string>")
	println("pretty-json -s <absolute-soure-file-path>")
	println("")
}
