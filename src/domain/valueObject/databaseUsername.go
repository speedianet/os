package valueObject

import (
	"errors"
	"regexp"
)

const databaseUsernameRegex string = `^[0-9a-zA-Z_-]+$`

type DatabaseUsername string

func NewDatabaseUsername(value string) (DatabaseUsername, error) {
	dbUser := DatabaseUsername(value)
	if !dbUser.isValid() {
		return "", errors.New("InvalidDatabaseUsername")
	}
	return dbUser, nil
}

func NewDatabaseUsernamePanic(value string) DatabaseUsername {
	dbUser := DatabaseUsername(value)
	if !dbUser.isValid() {
		panic("InvalidDatabaseUsername")
	}
	return dbUser
}

func (dbUser DatabaseUsername) isValid() bool {
	re := regexp.MustCompile(databaseUsernameRegex)
	return re.MatchString(string(dbUser))
}

func (dbUser DatabaseUsername) String() string {
	return string(dbUser)
}
