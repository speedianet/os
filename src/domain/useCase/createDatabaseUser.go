package useCase

import (
	"errors"
	"log"

	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/repository"
	"github.com/speedianet/os/src/domain/valueObject"
)

func CreateDatabaseUser(
	dbQueryRepo repository.DatabaseQueryRepo,
	dbCmdRepo repository.DatabaseCmdRepo,
	addDatabaseUser dto.CreateDatabaseUser,
) error {
	_, err := dbQueryRepo.GetByName(addDatabaseUser.DatabaseName)
	if err != nil {
		return errors.New("DatabaseNotFound")
	}

	if len(addDatabaseUser.Privileges) == 0 {
		addDatabaseUser.Privileges = []valueObject.DatabasePrivilege{
			valueObject.NewDatabasePrivilegePanic("ALL"),
		}
	}

	err = dbCmdRepo.AddUser(addDatabaseUser)
	if err != nil {
		return errors.New("CreateDatabaseUserError")
	}

	log.Printf(
		"Database user '%s' for '%s' created.",
		addDatabaseUser.Username.String(),
		addDatabaseUser.DatabaseName.String(),
	)

	return nil
}