package controllers

import (
	"net/http"
	"ori_saude_api/src/core/entities"
	"ori_saude_api/src/core/usecases"

	"github.com/gin-gonic/gin"
)

func CreateProfessional(c *gin.Context, uc *usecases.CreateProfessionalUsecase) {
	var professional entities.Professional
	if err := c.ShouldBindJSON(&professional); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.Execute(&professional); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, professional)
}

func GetProfessionals(c *gin.Context, uc *usecases.GetProfessionalsUsecase) {
	professionals, err := uc.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": professionals})
}

func GetProfessionalByID(c *gin.Context, uc *usecases.GetProfessionalByIDUsecase) {
	id := c.Param("id")
	prof, err := uc.Execute(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Professional not found"})
		return
	}
	c.JSON(200, prof)
}

// func UpdateProfessional(c *gin.Context) {
// 	id := c.Param("id")
// 	var professional entities.Professional
// 	if err := c.ShouldBindJSON(&professional); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	updatedProfessional := usecases.UpdateProfessional(id, &professional)
// 	if updatedProfessional == nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Professional not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, updatedProfessional)
// }

// func DeleteProfessional(c *gin.Context) {
// 	id := c.Param("id")
// 	if err := usecases.DeleteProfessional(id); err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Professional not found"})
// 		return
// 	}
// 	c.Status(http.StatusNoContent)
// }
