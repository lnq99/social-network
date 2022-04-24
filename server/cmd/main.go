package main

import (
	"log"
	"os"

	"app/config"
	"app/internal/controller"
	v1 "app/internal/controller/v1"
	"app/internal/driver"
	"app/internal/repository"
	"app/internal/service"
)

// Точка входа в программу.
func main() {
	confFile := ".env"
	if len(os.Args) == 2 {
		confFile = os.Args[1]
		log.Println(confFile)
	}

	var err error
	conf, err := config.LoadConfig(".", confFile)
	if err != nil {
		panic(err)
	}

	db := driver.Connect(conf.DbDriver, conf.DbHost, conf.DbPort, conf.DbUser, conf.DbPassword, conf.DbName)
	err = db.SQL.Ping()
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepo(db.SQL)
	services := service.GetServices(repo)

	ctrl := v1.NewController(services, &conf)

	router := controller.NewRouter()
	router = ctrl.SetupRouter(router)

	addr := conf.Host + ":" + conf.Port
	router.Run(addr)
}
