package controllers

import (
	"net/http"
	"ori_saude_api/src/core/usecases"

	"github.com/gin-gonic/gin"
)

func GetProfessionals(c *gin.Context) {
	c.JSON(http.StatusOK, usecases.GetProfessionals())
}
