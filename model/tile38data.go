package model

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/tidwall/gjson"
	"strconv"
)

type Tile38Data struct {
	Object []Tile38Object `json:"data"`
}

func FromScan(client *redis.Client, name string) (Tile38Data, error) {
	var tile38Data Tile38Data
	var err error
	v, err := client.Do("SCAN", name).Result()
	if err == nil {
		data, err := json.Marshal(v)
		if err == nil {
			jsonConverted := `{"data":` + string(data) + `}`
			lengthData := gjson.Get(jsonConverted, "data.1.#")
			if lengthData.Int() != 0 {
				var x int64 = 0
				for x = 0; x < lengthData.Int(); x++ {
					name := "data.1." + strconv.FormatInt(int64(x), 10)
					idName := name + ".0"
					contentName := name + ".1"
					id := gjson.Get(jsonConverted, idName)
					content := gjson.Get(jsonConverted, contentName)
					var tile38Object Tile38Object
					var tile38SubObject Tile38SubObject
					err = json.Unmarshal([]byte(content.String()), &tile38SubObject)
					if err == nil {
						tile38Object.Id = id.String()
						tile38Object.Object = tile38SubObject
						tile38Data.Object = append(tile38Data.Object, tile38Object)
					}
				}
			}
		}
	}
	return tile38Data, err
}

func (tile38Data Tile38Data) ToJsonString() (string, error) {
	var err error
	data, err := json.Marshal(tile38Data)
	return string(data), err
}
