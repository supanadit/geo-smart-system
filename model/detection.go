package model

// Trigger Type
const (
	inside  string = "inside"  // when an object is inside the specified area
	outside string = "outside" // when an object is outside the specified area
	enter   string = "enter"   // when an object that was not previously in the fence has entered the area
	exit    string = "exit"    // when an object that was previously in the fence has exited the area
	cross   string = "cross"   // when an object that was not previously in the fence has entered and exited the area
)

type Detection struct {
	Type        string   `json:"type"`
	Lat         string   `json:"lat"`
	Lng         string   `json:"lng"`
	TriggerType []string `json:"trigger"`
}
