package db

import (
	logging "Uploader/pkg"
	"fmt"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnection() (*gorm.DB, error) {
	log := logging.GetLogger()

	host := "localhost"
	port := "5432"
	user := "humo"
	password := "pass"
	dbname := "humo_db"

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dushanbe",
		host, user, password, dbname, port)

	conn, err := gorm.Open(postgresDriver.Open(connString))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Success connection to", host)
	return conn, nil
}
