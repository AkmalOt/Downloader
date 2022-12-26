package db

import (
	"Uploader/config"
	logging "Uploader/pkg"
	"fmt"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnection() (*gorm.DB, error) {
	log := logging.GetLogger()

	DataBase, err := config.ReturnDB()
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dushanbe",
		DataBase.Host, DataBase.User, DataBase.Password, DataBase.Dbname, DataBase.Port)

	conn, err := gorm.Open(postgresDriver.Open(connString))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Success connection to", DataBase.Host)
	return conn, nil
}
