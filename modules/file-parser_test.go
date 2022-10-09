package modules

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestGetAllDirsForLocation(t *testing.T) {
	path := "/Users/danieltuttle/projects/golang_projects/dpt_cli/test_resources/test_line_search"
	res, _ := GetAllFilesForLocation(path)

	res1 := "/Users/danieltuttle/projects/golang_projects/dpt_cli/test_resources/test_line_search/dir1/file2.txt"
	res2 := "/Users/danieltuttle/projects/golang_projects/dpt_cli/test_resources/test_line_search/dir1/dir2/file3.txt"
	res3 := "/Users/danieltuttle/projects/golang_projects/dpt_cli/test_resources/test_line_search/dir3/file4.txt"
	res4 := "/Users/danieltuttle/projects/golang_projects/dpt_cli/test_resources/test_line_search/file.txt"
	res5 := "/Users/danieltuttle/projects/golang_projects/dpt_cli/test_resources/test_line_search/dir1/dir2/dir4/file5.txt"

	assert.True(t, contains(res, res1))
	assert.True(t, contains(res, res2))
	assert.True(t, contains(res, res3))
	assert.True(t, contains(res, res4))
	assert.True(t, contains(res, res5))

}

func TestGetLinesThatMatchFromFile(t *testing.T) {
	res, _ := GetLinesThatMatchFromFile("/Users/danieltuttle/projects/golang_projects/dpt_cli/test_resources/test_line_search/dir1/dir2/dir4/file5.txt", "file5")
	e := "\nfile5.txt : line number: 1\nhello from file5\n"
	assert.Equal(t, e, res[0])
}

func TestGetFileNameFromPath(t *testing.T) {
	path := "/Users/danieltuttle/projects/golang_projects/dpt_cli/test_resources/test_line_search/dir1/dir2/dir4/file5.txt"
	res := GetFileNameFromPath(path)
	assert.Equal(t, "file5.txt", res)
}

func TestTesting(t *testing.T) {
	x := "/Users/danieltuttle/Desktop/ex.json"
	s := GetPrettyJson(x)
	WriteFileWithContents(s, "/Users/danieltuttle/Desktop/ex2.json")
}

//func PrettyString(str string) (string, error) {
//	var prettyJSON bytes.Buffer
//	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
//		return "", err
//	}
//	return prettyJSON.String(), nil
//}

func contains(list []string, s string) bool {
	for _, e := range list {
		if e == s {
			return true
		}
	}
	return false
}
