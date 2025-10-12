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
		//TODO: api.GET("/:id", controllers.GetProfessionalByID)
		//TODO: api.POST("/", controllers.CreateProfessional)
		//TODO: api.PUT("/:id", controllers.UpdateProfessional)
		//TODO: api.DELETE("/:id", controllers.DeleteProfessional)
	}

	return r
}
