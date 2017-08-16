package conf

import (
	"errors"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/prometheus/common/log"
	gcfg "gopkg.in/gcfg.v1"
)

// InitConfiguration : Read configuration file for Ganeca from external file.
func InitConfiguration(filename string) error {
	log.Info("Init ganeca with config file from :", filename)

	err := gcfg.ReadFileInto(&Configuration, filename)
	if err != nil {
		log.Fatal("Init configuration fail :", err.Error())
		return err
	}
	return nil
}

// InitConnection : Preparing connection for all.
func InitConnection() error {

	if Configuration.IsNull() {
		return errors.New("Configuration is not valid!")
	}

	// Prepare connection for master database
	if Configuration.Database.MASTERDB != "" {
		masterDB, err := sqlx.Open("postgres", Configuration.Database.MASTERDB)
		if err != nil {
			log.Fatalln("Error opening master database :", err.Error())
			return err
		}
		Connection.MASTERDB = masterDB
		_, err = Connection.MASTERDB.Query("SELECT 1")
		if err != nil {
			log.Fatalln("Error accessing master database :", err.Error())
			return err
		}
		log.Info("Connected to master database")
	} else {
		log.Error("Master database connection fail! Empty config")
	}

	// Prepare connection for slave database
	if Configuration.Database.SLAVEDB != "" {
		slaveDB, err := sqlx.Open("postgres", Configuration.Database.SLAVEDB)
		if err != nil {
			log.Fatalln("Error opening slave database :", err.Error())
			return err
		}
		Connection.SLAVEDB = slaveDB
		_, err = Connection.SLAVEDB.Query("SELECT 1")
		if err != nil {
			log.Fatalln("Error accessing slave database :", err.Error())
			return err
		}
	} else {
		log.Error("Slave database connection fail! Empty config")
	}

	// Prepare connection for redis pool
	if Configuration.Database.REDIS != "" {
		redisPool := &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", Configuration.Database.REDIS) },
		}
		Connection.REDIS = redisPool
	} else {
		log.Error("Redis connection fail! Empty config")
	}

	return nil
}
