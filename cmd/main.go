package main

import (
	"fmt"
	"log"
	"os"
	"parser/config"
	"parser/internal/repository"
	"parser/internal/service"
)

func main() {
	var cfgPath string

	switch len(os.Args[1:]) {
	case 1:
		cfgPath = os.Args[1]
	case 0:
		cfgPath = "./.env"
	default:
		log.Print("USAGE: go run [CONFIG_PATH]")
		return
	}

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Print(err)
		return
	}

	db, err := repository.LoadDB(cfg.Db.DriverName, cfg.Db.DataSourceName)
	if err != nil {
		log.Print(err)
		return
	}

	r := repository.NewRepository(db)

	s := service.NewService(r, cfg.Url)

	_, err = s.Samruk.GetLotsById("962298")
	if err != nil {
		fmt.Println(err)
		return
	}
}
