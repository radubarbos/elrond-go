package arwen

import (
	"io/ioutil"
	"path"
)

// GetBytecode -
func GetBytecode(relativePath string) ([]byte, error) {
	return ioutil.ReadFile(path.Join("../testdata", relativePath))
}
