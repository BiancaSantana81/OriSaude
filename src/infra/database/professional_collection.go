package database

import (
	"context"
	"log"
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

	// Push cria uma referência com ID automático
	newRef, err := ref.Push(r.ctx, nil)
	if err != nil {
		return err
	}

	prof.ID = newRef.Key // ID gerado pelo Firebase

	// Salva o objeto completo no Firebase
	if err := newRef.Set(r.ctx, prof); err != nil {
		return err
	}

	log.Printf("✅ Profissional criado em: professionals/%s", newRef.Key)
	return nil
}

// func GetAllProfessionals() []entities.Professional {
// 	return professionals
// }

// func GetProfessionalByID(id int) (*entities.Professional, bool) {
// 	for _, prof := range professionals {
// 		if prof.ID == id {
// 			return &prof, true
// 		}
// 	}
// 	return nil, false
// }

// func CreateProfessional(professional *entities.Professional) error {
// 	professional.ID = len(professionals) + 1
// 	professionals = append(professionals, *professional)
// 	return nil
// }

// func UpdateProfessional(professional *entities.Professional) error {
// 	for i, prof := range professionals {
// 		if prof.ID == professional.ID {
// 			professionals[i] = *professional
// 			return nil
// 		}
// 	}
// 	return nil
// }

// func DeleteProfessional(id int) error {
// 	for i, prof := range professionals {
// 		if prof.ID == id {
// 			professionals = append(professionals[:i], professionals[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return nil
// }