package usecases

import "ori_saude_api/src/core/entities"
import "ori_saude_api/src/core/repositories"

type UpdateProfessionalUsecase struct {
	repo repositories.ProfessionalRepository
}

func NewUpdateProfessionalUsecase(repo repositories.ProfessionalRepository) *UpdateProfessionalUsecase {
	return &UpdateProfessionalUsecase{repo: repo}
}

func (uc *UpdateProfessionalUsecase) Execute(prof *entities.Professional) error {
	return uc.repo.UpdateProfessional(prof)
}
