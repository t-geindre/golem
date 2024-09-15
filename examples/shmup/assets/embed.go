package assets

import (
	"embed"
	_ "embed"
)

//go:embed sheet.png
var Sheet []byte

//go:embed *.png
var Assets embed.FS
