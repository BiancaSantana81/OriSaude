package main

import (
	"ori_saude_api/src/infra/database"
	"ori_saude_api/src/infra/server"
)

func main() {
	database.InitFirebase()
	r := server.NewServer()
	r.Run(":8080")
}
