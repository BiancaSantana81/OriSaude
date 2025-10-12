package server

import (
	"context"
	"ori_saude_api/src/app/controllers"
	"ori_saude_api/src/infra/database"
	"ori_saude_api/src/core/usecases"

	"firebase.google.com/go/db"
	"github.com/gin-gonic/gin"
)

// func NewServer() *gin.Engine {
// 	r := gin.Default()

// 	api := r.Group("/professionals")
// 	{
// 		api.POST("/", controllers.CreateProfessional)
// 		// api.GET("/", controllers.GetProfessionals)
// 		// api.GET("/:id", controllers.GetProfessionalByID)
// 		// api.PUT("/:id", controllers.UpdateProfessional)
// 		// api.DELETE("/:id", controllers.DeleteProfessional)
// 	}
// 	return r
// }


func NewServer(client *db.Client, ctx context.Context) *gin.Engine {
	r := gin.Default()

	// Inicializa o repositório e o usecase
	profRepo := database.NewFirebaseProfessionalRepository(client, ctx)
	createProfessionalUC := usecases.NewCreateProfessionalUsecase(profRepo)

	api := r.Group("/professionals")
	{
		// Envelopando o controller para passar o usecase
		api.POST("/", func(c *gin.Context) {
			controllers.CreateProfessional(c, createProfessionalUC)
		})
	}

	return r
}
