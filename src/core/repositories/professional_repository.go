package repositories

import "ori_saude_api/src/core/entities"

type ProfessionalRepository interface {
	CreateProfessional(prof *entities.Professional) error
	GetAllProfessionals() ([]entities.Professional, error)
	GetProfessionalByID(id string) (*entities.Professional, error)
	// UpdateProfessional(prof *entities.Professional) error
	// DeleteProfessional(id int) error
}
