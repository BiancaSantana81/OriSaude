package database

import "ori_saude_api/src/core/entities"

func GetAllProfessionals() []entities.Professional {
	return professionals
}

func GetProfessionalByID(id int) (*entities.Professional, bool) {
	for _, prof := range professionals {
		if prof.ID == id {
			return &prof, true
		}
	}
	return nil, false
}

func CreateProfessional(professional *entities.Professional) error {
	professional.ID = len(professionals) + 1
	professionals = append(professionals, *professional)
	return nil
}

func UpdateProfessional(professional *entities.Professional) error {
	for i, prof := range professionals {
		if prof.ID == professional.ID {
			professionals[i] = *professional
			return nil
		}
	}
	return nil
}

func DeleteProfessional(id int) error {
	for i, prof := range professionals {
		if prof.ID == id {
			professionals = append(professionals[:i], professionals[i+1:]...)
			return nil
		}
	}
	return nil
}