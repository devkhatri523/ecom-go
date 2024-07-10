package database

import (
	"database/sql"
	"fmt"
	"github.com/ecom-go/config/config"
	"time"
)

type PostgresDatabase struct {
	db *sql.DB
}

func OpenPostgresSqlDatabase(host string, port int, username string, password string, database string) (*PostgresDatabase, error) {
	postgresDb := &PostgresDatabase{
		nil,
	}
	connMaxLifeTime := config.Default().GetInt("db.postgressql.connMaxLifeTime")
	maxOpenConn := config.Default().GetInt("db.postgressql.maxOpenConn")
	maxIdleConn := config.Default().GetInt("db.postgressql.maxIdleConn")
	param := config.Default().GetString("db.postgressql.param")

	postgresDb.Open(Options{
		Host:            host,
		Port:            port,
		Username:        username,
		Password:        password,
		Database:        database,
		Protocol:        "tcp",
		ConnMaxLifeTime: time.Duration(connMaxLifeTime),
		MaxOpenConn:     maxOpenConn,
		MaxIdleConn:     maxIdleConn,
		PARAM:           param,
	})
	return postgresDb, nil
}

func (m *PostgresDatabase) Open(options Options) {
	dsn, er := BuildDns(options)
	if er != nil {
		panic(er)
	}
	fmt.Sprintf("Opening database connection on host : %s port : %d database : %s username : %s",
		options.Host, options.Port, options.Database, options.Username)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Sprintf("Configuring database connection pool with ConnMaxLifeTime : %d MaxOpenConn : %d MaxIdleConn :  %d",
		options.ConnMaxLifeTime, options.MaxOpenConn, options.MaxIdleConn)
	db.SetConnMaxLifetime(options.ConnMaxLifeTime)
	db.SetMaxOpenConns(options.MaxOpenConn)
	db.SetMaxIdleConns(options.MaxIdleConn)
	m.db = db
	fmt.Sprintf("Database connection opened on host : %s port : %d database : %s username : %s",
		options.Host, options.Port, options.Database, options.Username)
}
func (m *PostgresDatabase) Close() {
	if m.db != nil {
		err := m.db.Close()
		if err != nil {
			fmt.Sprintf("Error while closing database connection %s", err)
			return
		}
	}
	fmt.Sprintf("Database connection closed.")
}
func (m *PostgresDatabase) Get() interface{} {
	if m.db == nil {
		panic("Database connection not initiated. Please call Open()")
	}
	return m.db
}
func (m *PostgresDatabase) Ping() error {
	if m.db == nil {
		panic("Database connection not initiated. Please call Open()")
	}
	err := m.db.Ping()
	if err != nil {
		return err
	} else {
		fmt.Sprintf("Postgressql database connection succeeded.")
		return nil
	}
}
