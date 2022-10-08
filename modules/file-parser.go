package modules

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Returns all files in path directory and subdirectories
func GetAllFilesForLocation(ap string) ([]string, error) {
	response := []string{}

	dirs := []string{}
	dirs = append(dirs, ap)

	for len(dirs) > 0 {
		currPath := dirs[0]
		dirs = dirs[1:]
		files, err := ioutil.ReadDir(currPath)
		if err == nil {
			for _, e := range files {
				if e.IsDir() {
					new_p := currPath + "/" + e.Name()
					dirs = append(dirs, new_p)
				} else {
					new_p := currPath + "/" + e.Name()
					response = append(response, new_p)
				}
			}
		} else {
			return nil, errors.New(fmt.Sprintf("There was a problem reading a file with path: %s ", currPath))
		}
	}
	return response, nil
}

func GetLinesThatMatchFromFile(path string, substr string) ([]string, error) {
	dat, _ := os.ReadFile(path)
	list := strings.Split(string(dat), "\\n")
	fileName := GetFileNameFromPath(path)

	lineNumber := 1
	res := []string{}
	for _, e := range list {
		if strings.Contains(e, substr) {
			newS := "\n" + fileName + " : line number: " + fmt.Sprint(lineNumber) + "\n" + e + "\n"
			res = append(res, newS)
		}
	}

	return res, nil
}

func GetFileNameFromPath(path string) string {
	res := strings.Split(path, "/")
	return res[len(res)-1]
}
