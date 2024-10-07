package assets

import (
	"embed"
	_ "embed"
)

//go:embed *.png *.otf *.xml
var Embedded embed.FS
