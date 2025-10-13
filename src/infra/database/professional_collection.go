package database

import (
	"context"
	"errors"
	"ori_saude_api/src/core/entities"
	"ori_saude_api/src/core/repositories"

	"firebase.google.com/go/db"
)

type FirebaseProfessionalRepository struct {
	client *db.Client
	ctx    context.Context
}

func NewFirebaseProfessionalRepository(client *db.Client, ctx context.Context) repositories.ProfessionalRepository {
	return &FirebaseProfessionalRepository{
		client: client,
		ctx:    ctx,
	}
}

func (r *FirebaseProfessionalRepository) CreateProfessional(prof *entities.Professional) error {
	ref := r.client.NewRef("professionals")

	newRef, err := ref.Push(r.ctx, nil)
	if err != nil {
		return err
	}

	prof.ID = newRef.Key

	if err := newRef.Set(r.ctx, prof); err != nil {
		return err
	}
	return nil
}

func (r *FirebaseProfessionalRepository) GetAllProfessionals() ([]entities.Professional, error) {
	var result map[string]entities.Professional

	err := r.client.NewRef("professionals").Get(r.ctx, &result)
	if err != nil {
		return nil, err
	}

	professionals := make([]entities.Professional, 0, len(result))
	for key, p := range result {
		p.ID = key
		professionals = append(professionals, p)
	}

	return professionals, nil
}

func (r *FirebaseProfessionalRepository) GetProfessionalByID(id string) (*entities.Professional, error) {
	var prof entities.Professional
	err := r.client.NewRef("professionals/"+id).Get(r.ctx, &prof)
	if err != nil {
		return nil, err
	}
	prof.ID = id
	return &prof, nil
}

func (r *FirebaseProfessionalRepository) UpdateProfessional(prof *entities.Professional) error {
	if prof.ID == "" {
		err := errors.New("professional ID is empty")
		return err
	}

	ref := r.client.NewRef("professionals/" + prof.ID)
	if err := ref.Set(r.ctx, prof); err != nil {
		return err
	}
	return nil
}

func (r *FirebaseProfessionalRepository) DeleteProfessional(id string) error {
	ref := r.client.NewRef("professionals/" + id)
	if err := ref.Delete(r.ctx); err != nil {
		return err
	}
	return nil
}
