package system

import (
	"fmt"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/rs/xid"
	"github.com/supanadit/geo-smart-system/model"
	"github.com/supanadit/geo-smart-system/model/tile38"
	"log"
	"net/http"
	"strings"
)

func Router(r *gin.Engine, client *redis.Client) {
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
		hookURL := GetTile38HookURL(hookID)
		trigger := strings.Join(detection.TriggerType, ",")
		var status map[string]interface{}
		var httpStatus int
		if trigger == "" {
			status = gin.H{"status": "Please include trigger type such as 'enter','cross','exit','inside' or 'outside'"}
			httpStatus = http.StatusBadRequest
		} else {
			client.Do("SETHOOK", hookID, hookURL, "NEARBY", detection.Type, "FENCE", "DETECT", trigger, "COMMANDS", "set", "POINT", detection.Lat, detection.Lng)
			if err != nil {
				status = gin.H{"status": "Unknown Error"}
				httpStatus = http.StatusInternalServerError
			} else {
				status = gin.H{"status": "Ok"}
				httpStatus = http.StatusOK
			}
		}
		c.Writer.Header().Set("Content-Type", "application/json")
		c.JSON(httpStatus, status)
	})

	r.GET("/detection/call", func(c *gin.Context) {
		hookID := c.Query("hook")
		if hookID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "Wrong Request"})
		} else {
			client.Do("DELHOOK", hookID)
			c.JSON(http.StatusOK, gin.H{"status": "OK"})
		}
	})

	r.GET("/ws", func(c *gin.Context) {
		websocketHandler(c.Writer, c.Request)
	})

	r.Static("/public", "./public")
	r.Static("/assets", "./assets")
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
