package usecases

import (
	"ori_saude_api/src/core/entities"
	"ori_saude_api/src/infra/database"
	"strconv"
)

func UpdateProfessional(id string, updatedData *entities.Professional) *entities.Professional {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}
	professional, found := database.GetProfessionalByID(intID)
	if !found || professional == nil {
		return nil
	}

	if updatedData.Name != "" {
		professional.Name = updatedData.Name
	}
	if updatedData.Category != "" {
		professional.Category = updatedData.Category
	}
	if updatedData.City != "" {
		professional.City = updatedData.City
	}
	if updatedData.Contact != "" {
		professional.Contact = updatedData.Contact
	}

	if err := database.UpdateProfessional(professional); err != nil {
		return nil
	}
	return professional
}