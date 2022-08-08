package zip

import (
	"os"
	"testing"
)

const tempdir = "../tempdir/"

func TestZip(t *testing.T) {
	os.Mkdir(tempdir, 0755)
	Zip(tempdir+"testzip.zip", "../zip")
}

func TestUnzip(t *testing.T) {
	Unzip(tempdir+"testzip.zip", tempdir)
}
