package main

import (
	"log"
	"ori_saude_api/src/infra/database"
	"ori_saude_api/src/infra/server"
)

func main() {
	// Inicializa Firebase
	database.InitFirebase() // isso configura database.Client e database.Ctx

	if database.Client == nil || database.Ctx == nil {
		log.Fatal("Firebase client or context is nil")
	}

	// Passa para o servidor
	r := server.NewServer(database.Client, database.Ctx)

	r.Run(":8080")
}


