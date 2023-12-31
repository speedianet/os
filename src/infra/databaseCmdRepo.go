package infra

import (
	"errors"

	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/valueObject"
	databaseInfra "github.com/speedianet/os/src/infra/database"
)

type DatabaseCmdRepo struct {
	dbType valueObject.DatabaseType
}

func NewDatabaseCmdRepo(
	dbType valueObject.DatabaseType,
) *DatabaseCmdRepo {
	return &DatabaseCmdRepo{
		dbType: dbType,
	}
}

func (repo DatabaseCmdRepo) Add(dbName valueObject.DatabaseName) error {
	switch repo.dbType {
	case "mysql":
		return databaseInfra.MysqlDatabaseCmdRepo{}.Add(dbName)
	case "postgres":
		return databaseInfra.PostgresDatabaseCmdRepo{}.Add(dbName)
	default:
		return errors.New("DatabaseTypeNotSupported")
	}
}

func (repo DatabaseCmdRepo) Delete(dbName valueObject.DatabaseName) error {
	switch repo.dbType {
	case "mysql":
		return databaseInfra.MysqlDatabaseCmdRepo{}.Delete(dbName)
	case "postgres":
		return databaseInfra.PostgresDatabaseCmdRepo{}.Delete(dbName)
	default:
		return errors.New("DatabaseTypeNotSupported")
	}
}

func (repo DatabaseCmdRepo) AddUser(addDatabaseUser dto.AddDatabaseUser) error {
	switch repo.dbType {
	case "mysql":
		return databaseInfra.MysqlDatabaseCmdRepo{}.AddUser(addDatabaseUser)
	case "postgres":
		return databaseInfra.PostgresDatabaseCmdRepo{}.AddUser(addDatabaseUser)
	default:
		return errors.New("DatabaseTypeNotSupported")
	}
}

func (repo DatabaseCmdRepo) DeleteUser(
	dbName valueObject.DatabaseName,
	dbUser valueObject.DatabaseUsername,
) error {
	switch repo.dbType {
	case "mysql":
		return databaseInfra.MysqlDatabaseCmdRepo{}.DeleteUser(dbName, dbUser)
	case "postgres":
		return databaseInfra.PostgresDatabaseCmdRepo{}.DeleteUser(dbName, dbUser)
	default:
		return errors.New("DatabaseTypeNotSupported")
	}
}
