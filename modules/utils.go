package modules

import (
	"log"
	"os"
)

func WriteFileWithContents(content string, path string) {
	f, err := os.Create(path)
	if err != nil {
		if err != nil {
			log.Fatal(err)
		}
	}

	defer f.Close()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		log.Fatal(err2)
	}
}
