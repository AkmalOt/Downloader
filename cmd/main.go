package main

import (
	"Uploader/config"
	"Uploader/internal/db"
	"Uploader/internal/repository"
	"Uploader/internal/server"
	"Uploader/internal/services"
	"Uploader/pkg"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
)

func main() {

	err := execute()
	if err != nil {
		log.Println(err)
	}
}

func execute() error {
	log := logging.GetLogger()
	router := mux.NewRouter()

	connection, err := db.GetDbConnection()
	if err != nil {
		log.Fatal("error on DB-connection", err)
	}
	newRepository := repository.NewRepository(connection)

	newServices := services.NewServices(newRepository)
	//newServicesConf := services.NewServices()
	newServer := server.NewServer(router, newServices)

	newServer.Init()

	getConfig, err := config.GetConfig()
	if err != nil {
		log.Fatal("GetConfig is crashed", err)
	}
	address := net.JoinHostPort(getConfig.Host, getConfig.Port)
	srv := http.Server{
		Addr:    address,
		Handler: router,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Println("err in ListenAndServe", err)
	}
	return nil
}
