package usecases

import (
	"ori_saude_api/src/core/entities"
	"ori_saude_api/src/infra/database"
	"strconv"
)

func GetProfessionalByID(id string) *entities.Professional {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil
	}
	var professional *entities.Professional

	professional, found := database.GetProfessionalByID(intID)
	if !found || professional == nil {
		return nil
	}
	return professional
}
