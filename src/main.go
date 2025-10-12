package main

import (
	"log"
	"ori_saude_api/src/infra/database"
	"ori_saude_api/src/infra/server"
)

func main() {
	database.InitFirebase()

	if database.Client == nil || database.Ctx == nil {
		log.Fatal("Firebase client or context is nil")
	}

	r := server.NewServer(database.Client, database.Ctx)
	r.Run(":8080")
}


