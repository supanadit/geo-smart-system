package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/rs/xid"
	"github.com/supanadit/geo-smart-system/model"
	"github.com/supanadit/geo-smart-system/model/tile38"
	"log"
	"net/http"
)

func main() {
	port := "8080"

	client := redis.NewClient(&redis.Options{
		Addr: "192.168.99.100:9851",
	})

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/id/get/unique", func(c *gin.Context) {
		id := xid.New()
		c.JSON(200, gin.H{"id": id.String()})
	})

	r.POST("/point/set", func(c *gin.Context) {
		var location model.Location
		err := c.BindJSON(&location)
		client.Do("SET", location.Type, location.Id, "POINT", location.Lat, location.Lng)
		var status map[string]interface{}
		var httpStatus int
		if err != nil {
			status = gin.H{"status": "Unknown Error"}
			httpStatus = http.StatusInternalServerError
		} else {
			status = gin.H{"status": "Ok"}
			httpStatus = http.StatusOK
		}
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(httpStatus, status)
	})

	r.POST("/point/unset", func(c *gin.Context) {
		var location model.Location
		err := c.BindJSON(&location)
		client.Do("DEL", location.Type, location.Id)
		var status map[string]interface{}
		var httpStatus int
		if err != nil {
			status = gin.H{"status": "Unknown Error"}
			httpStatus = http.StatusInternalServerError
		} else {
			status = gin.H{"status": "Ok"}
			httpStatus = http.StatusOK
		}
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(httpStatus, status)
	})

	r.GET("/point/get", func(c *gin.Context) {
		t := c.DefaultQuery("type", "user")
		data, _ := tile38.FromScan(client, t)
		c.JSON(http.StatusOK, data)
	})

	r.GET("/point/get/stream", func(c *gin.Context) {
		w := c.Writer
		t := c.DefaultQuery("type", "user")
		data, _ := tile38.FromScan(client, t)

		_ = sse.Encode(w, sse.Event{
			Event: "message",
			Data:  data,
		})
	})

	r.POST("/detection/set", func(c *gin.Context) {
		var detection model.Detection
		err := c.BindJSON(&detection)
		hookID := "HOOK-" + xid.New().String()
		fmt.Printf("Set HOOK with ID : %s \n", hookID)
		hookURL := "http://192.168.99.1:" + port + "/detection/call"
		client.Do("SETHOOK", hookID, hookURL, "NEARBY", detection.Type, "FENCE", "DETECT", "enter", "COMMANDS", "set", "POINT", detection.Lat, detection.Lng)
		var status map[string]interface{}
		var httpStatus int
		if err != nil {
			status = gin.H{"status": "Unknown Error"}
			httpStatus = http.StatusInternalServerError
		} else {
			status = gin.H{"status": "Ok"}
			httpStatus = http.StatusOK
		}
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(httpStatus, status)
	})

	r.GET("/detection/call", func(c *gin.Context) {
		fmt.Println("Called")
		c.JSON(200, gin.H{"status": "test"})
	})

	r.GET("/ws", func(c *gin.Context) {
		websocketHandler(c.Writer, c.Request)
	})

	r.Static("/public", "./public")
	r.Static("/assets", "./assets")

	_ = r.Run(":" + port)
}

var websocketUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketUpgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("Received : %s \n", msg)
		_ = conn.WriteMessage(t, []byte("OK"))
	}
}
