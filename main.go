package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/supanadit/geo-smart-system/system"
	"golang.org/x/crypto/acme/autocert"
	"log"
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
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}
