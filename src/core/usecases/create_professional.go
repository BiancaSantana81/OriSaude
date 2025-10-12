package usecases

import (
	"ori_saude_api/src/core/entities"
	"ori_saude_api/src/core/repositories"
)

type CreateProfessionalUsecase struct {
	repo repositories.ProfessionalRepository
}

func NewCreateProfessionalUsecase(repo repositories.ProfessionalRepository) *CreateProfessionalUsecase {
	return &CreateProfessionalUsecase{repo: repo}
}

func (uc *CreateProfessionalUsecase) Execute(prof *entities.Professional) error {
	return uc.repo.CreateProfessional(prof)
}
