package usecases

import (
	"ori_saude_api/src/infra/database"
	"strconv"
)

func DeleteProfessional(id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return database.DeleteProfessional(intID)
}
