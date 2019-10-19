package model

type Tile38SubObject struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func (tile38SubObject Tile38SubObject) Lng() float64 {
	return tile38SubObject.Coordinates[0]
}

func (tile38SubObject Tile38SubObject) Lat() float64 {
	return tile38SubObject.Coordinates[1]
}
