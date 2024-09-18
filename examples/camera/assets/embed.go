package assets

import (
	"embed"
	_ "embed"
)

//go:embed *.png *.tmj *.tsx
var FS embed.FS

//go:embed map.tmj
var Map []byte
