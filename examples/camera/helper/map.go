package helper

import (
	"encoding/json"
)

type MapLayer struct {
	Name string `json:"name"`
	Data []int  `json:"data"`
}

type Map struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Layers []MapLayer
}

func LoadMap(scr []byte) *Map {
	m := &Map{}
	err := json.Unmarshal(scr, &m)

	if err != nil {
		panic(err)
	}

	return m
}
