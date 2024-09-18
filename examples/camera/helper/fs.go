package helper

import "github.com/t-geindre/golem/examples/camera/assets"

func readFile(path string) []byte {
	fBytes, err := assets.FS.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return fBytes
}
