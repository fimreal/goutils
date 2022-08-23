package file

import (
	"strconv"
	"testing"
)

func TestAppendToFile(t *testing.T) {
	filename := "../tempdir/appendtofile.txt"
	for i := 0; i < 10; i++ {
		b := []byte(strconv.Itoa(i) + " 你好\n")
		err := AppendToFile(filename, b)
		if err != nil {
			t.Fatal(err)
		}
	}
}
