package main

import (
	"dpt_cli/modules"
	"fmt"
	"log"
	"os"
)

func main() {
	var Reset = "\033[0m"
	var Blue = "\033[34m"
	var Green = "\033[32m"

	args := os.Args
	if len(args) > 1 {
		prog := args[1]
		if prog == "pretty-json" {
			if len(args) > 3 {
				firstFlag := args[2]
				firstValue := args[3]
				if firstFlag == "-s" && len(firstValue) > 0 {
					s, e := modules.GetPrettyJson(firstValue)
					if e == nil {
						modules.WriteFileWithContents(s, firstValue)
					} else {
						log.Fatal("There was an error processing the JSON file: ", e)
					}
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
					numStringInstances := 0
					if e == nil {
						for _, e := range res {
							res, e := modules.GetLinesThatMatchFromFile(e, secondValue)
							numStringInstances += len(res)
							if e == nil {
								for _, e := range res {
									print(Green + e + Reset)
								}
							}
						}
						println("")
						println(fmt.Sprintf(Blue+"There are %d instances of this string"+Reset, numStringInstances))
						println("")
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
