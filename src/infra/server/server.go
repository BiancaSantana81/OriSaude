package server

import (
	"github.com/gin-gonic/gin"
	"ori_saude_api/src/app/controllers"
)

func NewServer() *gin.Engine {
	r := gin.Default()

	api := r.Group("/professionals")
	{
		api.GET("/", controllers.GetProfessionals)
		api.GET("/:id", controllers.GetProfessionalByID)
		api.POST("/", controllers.CreateProfessional)
		api.PUT("/:id", controllers.UpdateProfessional)
		api.DELETE("/:id", controllers.DeleteProfessional)
	}

	return r
}
