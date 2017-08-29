package main

import (
	"fmt"
	"os"

	"github.com/fadhilimamk/ganeca/src/conf"
	"github.com/fadhilimamk/ganeca/src/log"
	"github.com/fadhilimamk/ganeca/src/news"
)

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
	// var data []students.Student

	// data = students.GetAllStudents()

	// for _, student := range data {
	// 	fmt.Println(student.ToString())
	// }

	news.Init()

	for i, item := range news.ItemData {
		fmt.Printf("%d\t%s\n", i, item.GetTitle())
		fmt.Printf("\t\t%s\n", item.GetImage())
		fmt.Printf("\t\t%d\n", item.GetDate())
		fmt.Printf("\t\t%s\n\n", item.GetDescription())
	}

	fmt.Println("finish")

}
