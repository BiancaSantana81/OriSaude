package usecases

import (
	"ori_saude_api/src/core/entities"
	"ori_saude_api/src/infra/database"
)

func GetProfessionals() []entities.Professional {
	return database.GetAllProfessionals()
}
