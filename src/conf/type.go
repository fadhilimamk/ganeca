package conf

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
)

// Configuration : configuration object, represented INI config file.
var Configuration configuration

// Connection : connection object to access data.
var Connection connection

type configuration struct {
	Server   server
	Database database
}

type server struct {
	PORT        string
	ENVIRONMENT string
	GINMODE     string
}

type database struct {
	MASTERDB string
	SLAVEDB  string
	REDIS    string
}

type connection struct {
	MASTERDB *sqlx.DB
	SLAVEDB  *sqlx.DB
	REDIS    *redis.Pool
}

func (c configuration) IsNull() bool {
	if (configuration{}) == c {
		return true
	}
	return false
}

func (c connection) IsNull() bool {
	if (connection{}) == c {
		return true
	}
	return false
}
