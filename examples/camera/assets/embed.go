package assets

import (
	"embed"
	_ "embed"
)

//go:embed *.png *.tmj *.tsx
var FS embed.FS

//go:embed tileset.png
var Tileset []byte

//go:embed map.tmj
var Map []byte
