package valueObject

import (
	"errors"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type ServiceName string

var NativeSvcNamesWithAliases = map[string][]string{
	"php":   {"lsphp", "php-fpm", "php-cgi", "litespeed", "openlitespeed"},
	"node":  {"nodejs"},
	"mysql": {"mysqld", "mariadb", "percona", "perconadb"},
	"redis": {"redis-server"},
}

func NewServiceName(value string) (ServiceName, error) {
	servicesName := maps.Keys(NativeSvcNamesWithAliases)
	if slices.Contains(servicesName, value) {
		return ServiceName(value), nil
	}

	for _, serviceName := range servicesName {
		if slices.Contains(
			NativeSvcNamesWithAliases[serviceName],
			value,
		) {
			return ServiceName(value), nil
		}
	}

	return "", errors.New("InvalidServiceName")
}

func NewServiceNamePanic(value string) ServiceName {
	sn, err := NewServiceName(value)
	if err != nil {
		panic(err)
	}
	return sn
}

func (sn ServiceName) String() string {
	return string(sn)
}
