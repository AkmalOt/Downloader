package server

import (
	"Uploader/internal/models"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"os"
)

func (s *Server) Registration(w http.ResponseWriter, r *http.Request) {

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var userInfo models.AuthInfo

	err = json.Unmarshal(bytes, &userInfo)
	if err != nil {
		log.Println(err)
		return
	}

	err = s.Services.Register(&userInfo)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(200)

}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var userInfo models.AuthInfo

	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		log.Println(err)
		return
	}

	token, err := s.Services.Login(&userInfo)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(token)
	w.Write([]byte(token))
	w.WriteHeader(200)
}

func (s *Server) FolderCreator(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	value := ctx.Value(userID)
	userId := value.(string)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var FolderInfo models.Folder

	FolderInfo.UserID = userId

	err = json.Unmarshal(body, &FolderInfo)
	if err != nil {
		log.Println(err)
		return
	}

	s.Services.FolderCreation(&FolderInfo)

	log.Println("Folder created successful")
}

func (s *Server) GetFoldersFromParent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	value := ctx.Value(userID)
	userId := value.(string)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var FolderInfo models.Folder

	FolderInfo.UserID = userId

	err = json.Unmarshal(body, &FolderInfo)
	if err != nil {
		log.Println(err)
		return
	}

	Folders, err := s.Services.GetFoldersFromParent(&FolderInfo)
	if err != nil {
		return
	}

	//for _, getUser := range Folders {
	//	log.Println("*", getUser)

	data, err := json.MarshalIndent(Folders, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(data))
	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *Server) GetParentFolders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	value := ctx.Value(userID)
	userId := value.(string)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var FolderInfo models.Folder
	//var FilesInfo models.File

	FolderInfo.UserID = userId
	//FolderInfo.FolderID = ""

	err = json.Unmarshal(body, &FolderInfo)
	if err != nil {
		log.Println(err)
		return
	}

	ParentId, Folders, err := s.Services.GetParentFolders(&FolderInfo)
	if err != nil {
		return
	}

	log.Println(ParentId)
	//FilesInfo.FolderID = ParentId
	//Files, err := s.Services.GetFiles(&FilesInfo)

	FolderData, err := json.MarshalIndent(Folders, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	//FileData, err := json.MarshalIndent(Files, "", "  ")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	log.Println(string(FolderData))
	_, err = w.Write(FolderData)
	if err != nil {
		log.Println(err)
		return
	}
	//_, err = w.Write(FileData)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	log.Println(FolderInfo.FolderID, "haha", ParentId)
	//log.Println(FilesInfo.FolderID, "haha", FileData)
}

func (s *Server) GetFiles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	value := ctx.Value(userID)
	userId := value.(string)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var FolderInfo models.Folder
	var FilesInfo models.File

	FolderInfo.UserID = userId
	FilesInfo.UserID = userId

	err = json.Unmarshal(body, &FilesInfo)
	if err != nil {
		log.Println(err)
		return
	}

	parentId, _, _ := s.Services.GetParentFolders(&FolderInfo)
	log.Println(parentId)

	FilesInfo.FolderID = parentId
	log.Println(FilesInfo, "test")

	Files, err := s.Services.GetFiles(&FilesInfo)
	log.Println(Files)
	FileData, err := json.MarshalIndent(Files, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}

	_, err = w.Write(FileData)
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *Server) UploadFile(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	value := ctx.Value(userID)
	userId := value.(string)

	//body, err := io.ReadAll(r.Body)
	//if err != nil {
	//	log.Println(err)
	//}
	var FileInfo models.File

	FileInfo.UserID = userId

	formValue := r.FormValue("data")

	err := json.Unmarshal([]byte(formValue), &FileInfo)
	if err != nil {
		err := errors.WithStack(err)
		log.Println(err)
		return
	}

	//err = json.Unmarshal(body, &FileInfo)
	//if err != nil {
	//	log.Println("error in unmarshal", err)
	//	return
	//}

	log.Println(FileInfo.Name, FileInfo.UserID, FileInfo.FolderID)

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Print("error in fromfile", err)
		return
	}

	filename := header.Filename
	log.Println(file, filename, FileInfo)
	uploadFile, err := s.Services.SaveFile(file, filename, &FileInfo)
	if err != nil {
		log.Print("error in saveimage", err)
		return
	}

	//uploadFile.TargetUrl = "123123"
	//uploadFile.FolderID = "47137b2a-7091-11ed-92eb-7c8ae16c8c64"
	log.Println(file, filename, FileInfo)
	err = s.Services.UploadFile(uploadFile)
	if err != nil {
		log.Print("error in upload file", err)
		return
	}

	w.WriteHeader(200)

}

func (s *Server) DownloadFile(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	value := ctx.Value(userID)
	userId := value.(string)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var FileInfo models.File

	FileInfo.UserID = userId

	err = json.Unmarshal(body, &FileInfo)
	if err != nil {
		log.Println(err)
		return
	}

	FileData, err := s.Services.DownloadFiles(FileInfo.ID)
	if err != nil {
		log.Println(err)
		return
	}
	Validator, err := s.Services.Repository.ValidationForDownload(FileData)

	if Validator != FileInfo.UserID {
		log.Println("Access denied!")
		w.WriteHeader(451)
		return
	}

	file, err := os.OpenFile("files/"+FileData.Name, os.O_CREATE|os.O_RDWR, 0777)

	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		return
	}
	//f, err := io.ReadAll(file)
	//if err != nil {
	//	log.Println(err)
	//}
	//test := os.WriteFile(FileData.Name, f, 666)
	//log.Println(test)
	//
	//w.WriteHeader(202)
}
