package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/rs/xid"
	"github.com/supanadit/geosmartsystem/model"
	"github.com/supanadit/geosmartsystem/model/tile38"
	"net/http"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:9851",
	})
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/unique/id", func(c *gin.Context) {
		id := xid.New()
		c.JSON(200, gin.H{"id": id.String()})
	})
	r.POST("/set/point", func(c *gin.Context) {
		var location model.Location
		_ = c.BindJSON(&location)
		c.Writer.Header().Set("Content-Type", "application/json")
		client.Do("SET", location.Type, location.Id, "POINT", location.Lat, location.Lng)
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.GET("/stream", func(c *gin.Context) {
		w := c.Writer
		t := c.DefaultQuery("type", "user")
		r := c.DefaultQuery("request", "")
		data, _ := tile38.FromScan(client, t)
		if r == "sse" {
			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")
			_ = sse.Encode(w, sse.Event{
				Event: "message",
				Data:  data,
			})
		} else {
			c.JSON(http.StatusOK, data)
		}
	})
	r.Static("/public", "./public")
	r.Static("/assets", "./assets")
	_ = r.Run()
}
