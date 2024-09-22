package assets

import (
	"embed"
	_ "embed"
)

//go:embed *.png *.otf
var Embedded embed.FS
