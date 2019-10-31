package model

type Location struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Lat  string `json:"lat"`
	Lng  string `json:"lng"`
}
