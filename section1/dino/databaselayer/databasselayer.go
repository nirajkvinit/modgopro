package databaselayer

import (
	"errors"
)

const (
	MYSQL uint8 = iota
	SQLITE
	MONGODB
	COCKROACHDB
)

type DinoDBHandler interface {
	GetAvailableDynos()
}

var DBTypeNotSupported = errors.New("The Database type provided is not suported...")

func GetDatabaseHandler(dbtype uint8) (DinoDBHandler, error) {
	switch dbtype {
	case MYSQL:
		return NewMySQLHandler(), nil
	case MONGODB:
		return NewMongoDBHandler(), nil
	case SQLITE:
		return NewSQLiteHandler(), nil
	}
	return nil, DBTypeNotSupported
}
