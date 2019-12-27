package model

type Trigger string

const (
	inside  Trigger = "inside"  // when an object is inside the specified area
	outside Trigger = "outside" // when an object is outside the specified area
	enter   Trigger = "enter"   // when an object that was not previously in the fence has entered the area
	exit    Trigger = "exit"    // when an object that was previously in the fence has exited the area
	cross   Trigger = "cross"   // when an object that was not previously in the fence has entered and exited the area
)

type Detection struct {
	Type        string    `json:"type"`
	Lat         string    `json:"lat"`
	Lng         string    `json:"lng"`
	TriggerType []Trigger `json:"trigger_type"`
}
