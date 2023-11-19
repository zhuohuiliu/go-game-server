package main

import "go-game-server/gnet"

func main() {
	server := gnet.NewServer("127.0.0.1", 8882)
	server.Serve()
}
