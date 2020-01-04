package tile38

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/supanadit/geo-smart-system/model"
	"strconv"
)

type SubObject struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func (tile38SubObject SubObject) Lng() float64 {
	return tile38SubObject.Coordinates[0]
}

func (tile38SubObject SubObject) Lat() float64 {
	return tile38SubObject.Coordinates[1]
}

func FromLocation(location model.Location) SubObject {
	var tile38SubObject SubObject
	tile38SubObject.Type = "Point"
	lng, _ := strconv.ParseFloat(location.Lng, 64)
	lat, _ := strconv.ParseFloat(location.Lat, 64)
	tile38SubObject.Coordinates = []float64{
		lng,
		lat,
	}
	return tile38SubObject
}

func GetDataLocation(client *redis.Client, typeLocation string, id string) (Object, error) {
	var tile38Object Object
	var err error
	data, err := client.Do("GET", typeLocation, id).Result()
	if err == nil {
		dataMarshal, err := json.Marshal(data)
		if err == nil {
			fmt.Println(string(dataMarshal))
			err = json.Unmarshal(dataMarshal, &tile38Object)
		}
	}
	return tile38Object, err
}
