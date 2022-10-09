package modules

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

func GetPrettyJson(path string) string {
	d, _ := ioutil.ReadFile(path)
	x, e := prettyString(string(d))
	if e != nil {
		println("There was an error")
	}
	return x
}

func prettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
