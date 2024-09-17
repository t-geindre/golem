package helper

type TDataSet struct {
	Columns int         `xml:"columns,attr"`
	Tw      int         `xml:"tilewidth,attr"`
	Th      int         `xml:"tileheight,attr"`
	Image   TDataSource `xml:"image"`
	Tiles   []TDataTile `xml:"tile"`
}

type TDataSource struct {
	Src string `xml:"source,attr"`
}

type TDataTile struct {
	Id     int          `xml:"id,attr"`
	Frames []TDataFrame `xml:"animation>frame"`
}

type TDataFrame struct {
	TileId   int `xml:"tileid,attr"`
	Duration int `xml:"duration,attr"`
}
