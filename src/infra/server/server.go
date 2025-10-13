package server

import (
	"context"
	"ori_saude_api/src/app/controllers"
	"ori_saude_api/src/core/usecases"
	"ori_saude_api/src/infra/database"

	"firebase.google.com/go/db"
	"github.com/gin-gonic/gin"
)

func NewServer(client *db.Client, ctx context.Context) *gin.Engine {
	r := gin.Default()

	profRepo := database.NewFirebaseProfessionalRepository(client, ctx)
	createProfessionalUC := usecases.NewCreateProfessionalUsecase(profRepo)
	getProfessionalsUC := usecases.NewGetProfessionalsUsecase(profRepo)
	getProfessionalByIDUC := usecases.NewGetProfessionalByIDUsecase(profRepo)
	updateProfessionalUC := usecases.NewUpdateProfessionalUsecase(profRepo)
	deleteProfessionalByIDUC := usecases.NewDeleteProfessionalByIDUsecase(profRepo)

	api := r.Group("/professionals")
	{
		api.POST("/", func(c *gin.Context) {
			controllers.CreateProfessional(c, createProfessionalUC)
		})

		api.GET("/", func(c *gin.Context) {
			controllers.GetProfessionals(c, getProfessionalsUC)
		})

		api.GET("/:id", func(c *gin.Context) {
			controllers.GetProfessionalByID(c, getProfessionalByIDUC)
		})

		api.PUT("/:id", func(c *gin.Context) {
			controllers.UpdateProfessional(c, updateProfessionalUC)
		})

		api.DELETE("/:id", func(c *gin.Context) {
			controllers.DeleteProfessionalByID(c, deleteProfessionalByIDUC)
		})
	}

	return r
}
