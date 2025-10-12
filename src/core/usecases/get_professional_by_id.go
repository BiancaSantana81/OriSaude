package usecases

import (
	"ori_saude_api/src/core/entities"
	"ori_saude_api/src/core/repositories"
)

type GetProfessionalByIDUsecase struct {
	repo repositories.ProfessionalRepository
}

func NewGetProfessionalByIDUsecase(repo repositories.ProfessionalRepository) *GetProfessionalByIDUsecase {
	return &GetProfessionalByIDUsecase{repo: repo}
}

func (uc *GetProfessionalByIDUsecase) Execute(id string) (*entities.Professional, error) {
	return uc.repo.GetProfessionalByID(id)
}
