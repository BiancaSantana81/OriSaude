package usecases

import (
	"ori_saude_api/src/core/entities"
	"ori_saude_api/src/infra/database"
)

func CreateProfessional(professional *entities.Professional) error {
	return database.CreateProfessional(professional)
}
