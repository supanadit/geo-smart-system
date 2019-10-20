package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	socketio "github.com/googollee/go-socket.io"
	"github.com/supanadit/geosmartsystem/model"
	"github.com/supanadit/geosmartsystem/model/tile38"
	"log"
)

func main() {
	// Tile38
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:9851",
	})

	router := gin.Default()
	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"http://localhost:4200"},
	//	AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowWebSockets:  true,
	//	AllowCredentials: true,
	//	AllowWildcard:    true,
	//}))
	router.Use(cors.Default())
	router.GET("/points", func(c *gin.Context) {
		data, _ := tile38.FromScan(client, "sales")
		c.JSON(200, data)
	})
	// Socket.IO Start
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("Connected:", s.ID())
		data, _ := tile38.FromScan(client, "sales")
		s.Emit("points", data)
		return nil
	})
	server.OnEvent("/", "set-points", func(s socketio.Conn, msg string) {
		var location model.Location
		err = json.Unmarshal([]byte(msg), &location)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(location)
			data, err := tile38.GetDataLocation(client, "sales", location.Id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(data)
			}

		}
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		_ = s.Close()
		return last
	})
	server.OnError("/", func(e error) {
		fmt.Println("Meet Error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		fmt.Println("Closed", msg)
	})
	router.GET("/socket.io/", gin.WrapH(server))
	router.POST("/socket.io/", gin.WrapH(server))
	router.Handle("WS", "/socket.io/", WebSocketIO(server))
	router.Handle("WSS", "/socket.io/", WebSocketIO(server))
	router.GET("/ws", func(c *gin.Context) {
		server.ServeHTTP(c.Writer, c.Request)
	})
	go server.Serve()
	defer server.Close()
	// End Socket.IO
	_ = router.Run()
}

func WebSocketIO(server *socketio.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		server.ServeHTTP(c.Writer, c.Request)
	}
}
