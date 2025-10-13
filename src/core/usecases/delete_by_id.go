package usecases

import "ori_saude_api/src/core/repositories"

type DeleteProfessionalByIDUsecase struct {
	repo repositories.ProfessionalRepository
}

func NewDeleteProfessionalByIDUsecase(repo repositories.ProfessionalRepository) *DeleteProfessionalByIDUsecase {
	return &DeleteProfessionalByIDUsecase{repo: repo}
}

func (uc *DeleteProfessionalByIDUsecase) Execute(id string) error {
	return uc.repo.DeleteProfessional(id)
}
