package main

import (
	"air-server/db"
	"air-server/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()

}
