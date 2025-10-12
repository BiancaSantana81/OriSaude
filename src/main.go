package main

import (
	"ori_saude_api/src/infra/server"
)

func main() {
	r := server.NewServer()
	r.Run(":8080")
}
