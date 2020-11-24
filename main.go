package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/supanadit/geo-smart-system/system"
)

func main() {
	// Create Connection with Tile 38
	client := redis.NewClient(&redis.Options{
		Addr: system.GetTile38ConnectionAddress(),
	})
	// Create Gin Engine
	r := gin.Default()
	r.Use(cors.Default())
	// Call Router
	system.Router(r, client)
	// Run Server
	_ = r.Run(fmt.Sprintf(":%s", system.ServerPort))
}
