package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/supanadit/geosmartsystem/model"
	"github.com/supanadit/geosmartsystem/model/tile38"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:9851",
	})
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/set-points", func(c *gin.Context) {
		var location model.Location
		_ = c.BindJSON(&location)
		c.Writer.Header().Set("Content-Type", "application/json")
		client.Do("SET", location.Type, location.Id, "POINT", location.Lat, location.Lng)
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.GET("/stream", func(c *gin.Context) {
		writer := c.Writer
		writer.Header().Set("Content-Type", "text/event-stream")
		writer.Header().Set("Cache-Control", "no-cache")
		writer.Header().Set("Connection", "keep-alive")
		data, _ := tile38.FromScan(client, "sales")
		_ = sse.Encode(writer, sse.Event{
			Event: "message",
			Data:  data,
		})
	})
	r.Static("/public", "./public")
	_ = r.Run()
}
