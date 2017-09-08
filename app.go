package main

import (
	"fmt"
	"os"

	"github.com/fadhilimamk/ganeca/src/agenda"

	"github.com/fadhilimamk/ganeca/src/conf"
	"github.com/fadhilimamk/ganeca/src/ganeca"
	"github.com/fadhilimamk/ganeca/src/log"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	log.InitLogger()

	filename := fmt.Sprintf("./file/config/config.%s.ini", env)
	if err := conf.InitConfiguration(filename); err != nil {
		log.Fatal("Error initializing ambalwarsa!")
	}
	log.Info("Config loaded")

	err := conf.InitConnection()
	if err != nil {
		log.Error("Error initiating connection :", err.Error())
	}

}

func main() {

	gin.SetMode(conf.Configuration.Server.GINMODE)
	router = gin.Default()

	log.Info("Preparing data")
	// news.Init()
	agenda.Init()

	log.Info("Ganeca is listening you on port ", conf.Configuration.Server.PORT)

	router.GET("/news", ganeca.ListNewsHandler)
	router.GET("/news/:id", ganeca.NewsDetailHandler)
	router.GET("events/", ganeca.ListEventsHandler)

	router.Run(conf.Configuration.Server.PORT)

}
