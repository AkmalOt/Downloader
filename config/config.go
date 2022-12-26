package config

import (
	"Uploader/internal/models"
	logging "Uploader/pkg"
	"encoding/json"
	"io"
	"os"
)

func GetConfig() (*models.Config, error) {
	log := logging.GetLogger()

	//чтение и дессериализация конфигов
	file, err := os.Open("./config/config.json")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var config models.Config

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &config, err
}

func Direction() string {
	Direction := "D:/Server"
	return Direction
}

var MySingingKey = []byte("TestIsRealHard")

func ReturnDB() (models.HumoDataBase, error) {
	var DB models.HumoDataBase
	DB.Host = "localhost"
	DB.Port = "5432"
	DB.User = "humo"
	DB.Password = "pass"
	DB.Dbname = "humo_db"

	return DB, nil
}
