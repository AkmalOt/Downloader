package server

import (
	"Uploader/internal/services"
	"github.com/gorilla/mux"
)

type Server struct {
	Mux      *mux.Router
	Services *services.Services
}

func NewServer(mux *mux.Router, services *services.Services) *Server {
	return &Server{
		Mux:      mux,
		Services: services,
	}
}

func (s *Server) Init() {
	//s.Mux.HandleFunc("/registration", s.Registration).Methods("POST")
	//s.Mux.HandleFunc("/login", s.Login).Methods("POST")

	authRoute := s.Mux.PathPrefix("/auth").Subrouter()

	authRoute.HandleFunc("/registration", s.Registration).Methods("POST")
	authRoute.HandleFunc("/login", s.Login).Methods("POST")

	TestRoute := s.Mux.PathPrefix("/test").Subrouter()
	TestRoute.Use(s.TokenValidator)
	TestRoute.HandleFunc("/create_folder", s.FolderCreator)
	TestRoute.HandleFunc("/show_folder", s.GetFoldersFromParent)
	TestRoute.HandleFunc("/show_parent_folder", s.GetParentFolders)
	TestRoute.HandleFunc("/upload_file", s.UploaFile)

	//authRoute.HandleFunc("/login", s.Test).Methods("POST")
	//s.Mux.HandleFunc("/test", s.Test)
	//s.Mux.Handle("/test", s.Middleware.TokenValidator(http.HandlerFunc(s.Test)))

	//s.Mux.Use(middleware.TokenValidator)

}

// Swagger

// -- вопросы:
// 1. Можно ли будет дропать файлы из инсомния
// 2. а скачка файлов?

// Заметки:
// 1. Надо будет создать список юзеров который выводится на консоль
// 2. Надо будет добавить выбор юзера для передачи прав на скачивания файлов.
// 3.

// --- uuid
// --- Refresh Token
