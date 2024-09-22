package helper

import (
	"github.com/t-geindre/golem/examples/scenes/assets"
	"io/fs"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	if path[0:7] == "embd://" {
		return assets.Embedded.ReadFile(path[7:])
	} else {
		return os.ReadFile(path)
	}
}

func OpenFile(path string) (fs.File, error) {
	if path[0:7] == "embd://" {
		return assets.Embedded.Open(path[7:])
	} else {
		return os.Open(path)
	}
}
