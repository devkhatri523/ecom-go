package database

import (
	"errors"
	"fmt"
	"time"
	"v01/utils"
)

type Options struct {
	Host            string
	Port            int
	Username        string
	Password        string
	Database        string
	Protocol        string
	ConnMaxLifeTime time.Duration
	MaxOpenConn     int
	MaxIdleConn     int
	PARAM           string
}

type Database interface {
	Open(options Options)
	Get() interface{}
	Close()
	Ping() error
}

func BuildDns(options Options) (string, error) {
	handleError := func(msg string) (string, error) {
		return "", errors.New(msg)
	}
	if utils.IsBlank(options.Username) {
		return handleError("user name cannot be empty")
	}
	if utils.IsBlank(options.Password) {
		return handleError("password cannot be empty")
	}
	if utils.IsBlank(options.Host) {
		return handleError("host name cannot be empty")
	}
	if options.Port <= 0 {
		return handleError("port cannot be 0 or negative")
	}
	if utils.IsBlank(options.Database) {
		return handleError("database cannot be empty")
	}
	var protocol string
	if utils.IsBlank(options.Protocol) {
		protocol = "tcp"
	} else {
		protocol = options.Protocol
	}
	//user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
	var param string
	if options.PARAM == "" {
		param = "parseTime=true"
	} else {
		param = options.PARAM
	}
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?%s", options.Username, options.Password, protocol,
		options.Host, options.Port, options.Database, param), nil

}
