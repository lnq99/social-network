package main

import (
	"app/config"
	v1 "app/internal/controller/v1"
	"app/internal/driver"
	"app/internal/repository"
	"app/internal/service"
)

func main() {
	var err error
	conf, err := config.LoadConfig(".", ".env")
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

	addr := conf.Host + ":" + conf.Port

	_, _ = ctrl, addr
}
