package database

import (
	"database/sql"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type OrmDB struct {
	OrmInstance *gorm.DB
	Database    Database
}

func OpenORMWithDatabase(database Database) (*OrmDB, error) {
	ormDB := OrmDB{}
	if database == nil {
		return nil, errors.New("database object is nil")
	}
	var err error
	d := database.Get().(*sql.DB)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: d,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	ormDB.OrmInstance = gormDB
	ormDB.Database = database
	return &ormDB, nil
}

func OpenORM(host string, port int, username string, password string, database string) (*OrmDB, error) {
	postgersDb, err := OpenPostgresSqlDatabase(host, port, username, password, database)
	if err != nil {
		return nil, err
	}
	err = postgersDb.Ping()
	if err != nil {
		return nil, err
	}
	return OpenORMWithDatabase(postgersDb)
}
