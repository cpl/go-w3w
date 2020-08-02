package w3w

import (
	"fmt"
)

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (c *Coordinates) String() string {
	return fmt.Sprintf("%f,%f", c.Lat, c.Lng)
}

type Square struct {
	SouthWest Coordinates `json:"southwest"`
	NorthEast Coordinates `json:"northeast"`
}

type ResponseError struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type ResponseConvert struct {
	Country      string      `json:"country"`
	Words        string      `json:"words"`
	Language     string      `json:"language"`
	Map          string      `json:"map"`
	NearestPlace string      `json:"nearestPlace"`
	Square       Square      `json:"square"`
	Coordinates  Coordinates `json:"coordinates"`
}
