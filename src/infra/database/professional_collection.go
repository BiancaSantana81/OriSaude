package database

import "ori_saude_api/src/core/entities"

func GetAllProfessionals() []entities.Professional {
	return professionals
}