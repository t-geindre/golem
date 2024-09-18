package helper

type MDataMap struct {
	Width    int `json:"width"`
	Height   int `json:"height"`
	Layers   []MDataLayer
	Tilesets []MDataTileset `json:"tilesets"`
	Tw       int            `json:"tilewidth"`
	Th       int            `json:"tileheight"`
}

type MDataLayer struct {
	Name string `json:"name"`
	Data []int  `json:"data"`
}

type MDataTileset struct {
	Source string `json:"source"`
}
