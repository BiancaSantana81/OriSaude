package usecases

import (
	"ori_saude_api/src/core/entities"
	"ori_saude_api/src/core/repositories"
)

type GetProfessionalsUsecase struct {
	repo repositories.ProfessionalRepository
}

func NewGetProfessionalsUsecase(repo repositories.ProfessionalRepository) *GetProfessionalsUsecase {
	return &GetProfessionalsUsecase{repo: repo}
}

func (uc *GetProfessionalsUsecase) Execute() ([]entities.Professional, error) {
	return uc.repo.GetAllProfessionals()
}
