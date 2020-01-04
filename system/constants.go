package system

import "fmt"

const (
	// Tile 38
	Tile38Address = "127.0.0.1"
	Tile38Port    = "9851"
	// Server Port
	ServerPort = "8080"
)

func GetTile38ConnectionAddress() string {
	return fmt.Sprintf("%s:%s", Tile38Address, Tile38Port)
}

func GetTile38HookURL(hookID string) string {
	return fmt.Sprintf("http://%s/detection/call?hook=%s", GetTile38ConnectionAddress(), hookID)
}
