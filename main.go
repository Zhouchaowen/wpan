package main

import (
	"flag"
	"fmt"
	"os"
	"wpan/api"
	"wpan/config"
	"wpan/db"
	"wpan/logger"
	"wpan/model"
	"wpan/utils"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "config/conf.yaml", "run config yaml")
}

func main() {
	flag.Parse()
	config.InitConfig(configPath)
	db.InitDB(config.ServerConfig.MySQL)

	if isExist, _ := utils.PathExists(config.ServerConfig.APP.Location); !isExist {
		fmt.Println("Save File Directory does not exist")
		os.Exit(1)
	}

	log, err := logger.New("STORAGE-API")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	db.DB.AutoMigrate(&model.User{})
	db.DB.AutoMigrate(&model.Share{})
	db.DB.AutoMigrate(&model.File{})
	db.DB.AutoMigrate(&model.FileStore{})
	db.DB.AutoMigrate(&model.Folder{})

	r := api.SetupRoute(log)
	if err := r.Run(fmt.Sprintf(":%d", config.ServerConfig.Server.HTTPPort)); err != nil {
		log.Fatalf("run error %s ", err.Error())
	}
}
