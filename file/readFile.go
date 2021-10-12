package file

import (
	"io/ioutil"
)

func GetFileContentByte(fileName string) ([]byte, error) {
	out, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []byte{}, err
	}
	return out, nil
}
