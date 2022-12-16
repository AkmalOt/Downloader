package server

import (
	"Uploader/internal/models"
	logging "Uploader/pkg"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
)

//

func (s *Server) Registration(w http.ResponseWriter, r *http.Request) {
	log := logging.GetLogger()

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
	log := logging.GetLogger()

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
	log := logging.GetLogger()

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
	log := logging.GetLogger()

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
	log := logging.GetLogger()

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

	//ctx := r.Context()
	//value := ctx.Value(userID)
	//userId := value.(string)
	//
	//body, err := io.ReadAll(r.Body)
	//if err != nil {
	//	log.Println(err)
	//}
	//var FolderInfo models.Folder
	//var FilesInfo models.File
	//var FileAndFolders models.FilesAndFolders
	//
	//FolderInfo.UserID = userId
	//FilesInfo.UserID = userId
	//
	////FolderInfo.FolderID = ""
	//
	//err = json.Unmarshal(body, &FolderInfo)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//if FilesInfo.FolderID == "" {
	//	ParentId, Folders, err := s.Services.GetParentFolders(&FolderInfo)
	//	if err != nil {
	//		return
	//	}
	//
	//	FilesInfo.FolderID = ParentId
	//	Files, err := s.Services.GetFiles(&FilesInfo)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//
	//	log.Println(Folders, Files)
	//
	//	FileAndFolders.Folder = FolderInfo
	//	FileAndFolders.File = FilesInfo
	//
	//	Folder, err := json.MarshalIndent(FileAndFolders.Folder, "", "  ")
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	File, err := json.MarshalIndent(FileAndFolders.Folder, "", "  ")
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//
	//	Comma := append(Folder, File...)
	//	log.Println("this is Comma!!!", string(Comma))
	//
	//	FileAndFolderByte, err := json.MarshalIndent(FileAndFolders, "", "  ")
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//
	//	_, err = w.Write(FileAndFolderByte)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	log.Println("test in empty", FileAndFolders)
	//} else {
	//	_, Folders, err := s.Services.GetParentFolders(&FolderInfo)
	//	if err != nil {
	//		return
	//	}
	//
	//	Files, err := s.Services.GetFiles(&FilesInfo)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//
	//	log.Println(Folders, Files)
	//
	//	FileAndFolders.Folder = FolderInfo
	//	FileAndFolders.File = FilesInfo
	//
	//	Folder, err := json.MarshalIndent(FileAndFolders.Folder, "", "  ")
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	File, err := json.MarshalIndent(FileAndFolders.Folder, "", "  ")
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//
	//	Comma := append(Folder, File...)
	//	log.Println("this is Comma!!!", Comma)
	//
	//	FileAndFolderByte, err := json.MarshalIndent(FileAndFolders, "", "  ")
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//
	//	_, err = w.Write(FileAndFolderByte)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	log.Println("test in not-empty", FileAndFolders)
	//	//log.Println(ParentId)
	//	////FilesInfo.FolderID = ParentId
	//	////Files, err := s.Services.GetFiles(&FilesInfo)
	//	//
	//	//FolderData, err := json.MarshalIndent(Folders, "", "  ")
	//	//if err != nil {
	//	//	log.Println(err)
	//	//	return
	//	//}
	//
	//	//FileData, err := json.MarshalIndent(Files, "", "  ")
	//	//if err != nil {
	//	//	log.Println(err)
	//	//	return
	//	//}
	//
	//	//log.Println(string(FolderData))
	//	//_, err = w.Write(FolderData)
	//	//if err != nil {
	//	//	log.Println(err)
	//	//	return
	//	//}
	//	//_, err = w.Write(FileData)
	//	//if err != nil {
	//	//	log.Println(err)
	//	//	return
	//	//}
	//
	//}
}

func (s *Server) GetFiles(w http.ResponseWriter, r *http.Request) {
	log := logging.GetLogger()

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
	if FilesInfo.FolderID == "" {
		parenFoldertId, _, _ := s.Services.GetParentFolders(&FolderInfo)
		log.Println(parenFoldertId)

		FilesInfo.FolderID = parenFoldertId
		log.Println(FilesInfo, "test")

		Files, err := s.Services.GetFiles(&FilesInfo)
		log.Println(Files)
		FileData, err := json.MarshalIndent(Files, "", "  ")
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("he;p!!!", FilesInfo.ID, FilesInfo.Name, FilesInfo.UserID, FilesInfo.FolderID)
		_, err = w.Write(FileData)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		log.Println(FilesInfo, "test")

		Files, err := s.Services.GetFiles(&FilesInfo)
		log.Println(Files)
		FileData, err := json.MarshalIndent(Files, "", "  ")
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("he;p!!!", FilesInfo.ID, FilesInfo.Name, FilesInfo.UserID, FilesInfo.FolderID)
		_, err = w.Write(FileData)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (s *Server) UploadFile(w http.ResponseWriter, r *http.Request) {
	log := logging.GetLogger()

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

	file, header, err := r.FormFile("files")
	if err != nil {
		log.Print("error in fromfile", err)
		return
	}

	filename := header.Filename
	log.Println(file, filename, FileInfo)
	uploadFile, err := s.Services.SaveFile(file, filename, &FileInfo)
	if err != nil {
		log.Print("error in saveFile", err)
		return
	}

	//uploadFile.TargetUrl = "123123"
	//uploadFile.FolderID = "47137b2a-7091-11ed-92eb-7c8ae16c8c64"
	log.Println(file, filename, FileInfo)
	err = s.Services.UploadFile(uploadFile)
	if err != nil {
		log.Print("error in upload files", err)
		return
	}

	w.WriteHeader(200)

}

func (s *Server) DownloadFile(w http.ResponseWriter, r *http.Request) {
	log := logging.GetLogger()

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
	//f, err := io.ReadAll(files)
	//if err != nil {
	//	log.Println(err)
	//}
	//test := os.WriteFile(FileData.Name, f, 666)
	//log.Println(test)
	//
	//w.WriteHeader(202)
}

func (s *Server) ChangeFileName(w http.ResponseWriter, r *http.Request) {
	log := logging.GetLogger()

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

	err = s.Services.ChangeFileName(&FileInfo)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(200)
}

func (s *Server) ChangeFolderName(w http.ResponseWriter, r *http.Request) {
	log := logging.GetLogger()

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

	err = s.Services.ChangeFolderName(&FolderInfo)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(200)
}

func (s *Server) DeleteFile(w http.ResponseWriter, r *http.Request) {
	log := logging.GetLogger()

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

	_, err = s.Services.GetFileInfoByID(&FileInfo)

	err = os.Remove("./files/" + FileInfo.Name)
	if err != nil {
		log.Println(err)
		return
	}

	err = s.Services.DeleteFile(&FileInfo)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(200)
}

func (s *Server) GiveAccess(w http.ResponseWriter, r *http.Request) {
	log := logging.GetLogger()

	ctx := r.Context()
	value := ctx.Value(userID)
	userId := value.(string)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var Access models.AccessTo

	Access.UserID = userId

	err = json.Unmarshal(body, &Access)
	if err != nil {
		log.Println(err)
		return
	}

	err = s.Services.GiveAccess(&Access)
	w.WriteHeader(200)

}

func (s *Server) GetAccessedFiles(w http.ResponseWriter, r *http.Request) {
	log := logging.GetLogger()

	ctx := r.Context()
	value := ctx.Value(userID)
	userId := value.(string)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var Access models.AccessTo

	Access.AccessedID = userId

	err = json.Unmarshal(body, &Access)
	if err != nil {
		log.Println(err)
		return
	}

	data, err := s.Services.GetAccessedFiles(&Access)

	FileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = w.Write(FileData)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(202)

}

func (s *Server) DownloadAccessedFiles(w http.ResponseWriter, r *http.Request) {
	log := logging.GetLogger()

	ctx := r.Context()
	value := ctx.Value(userID)
	userId := value.(string)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var FileInfo models.File
	var AccessInfo models.AccessTo

	FileInfo.UserID = userId

	err = json.Unmarshal(body, &FileInfo)
	if err != nil {
		log.Println(err)
		return
	}

	FileData, err := s.Services.DownloadAccessedFiles(&FileInfo)
	if err != nil {
		log.Println(err)
		return
	}
	AccessInfo.AccessedID = userId
	AccessInfo.FileId = FileInfo.ID

	Validator, err := s.Services.Repository.ValidationForAccessDownload(&AccessInfo)

	//log.Println("1", Validator.ID, "2", Validator.AccessedID, "3", Validator.UserID, "4", Validator.FileId)
	if Validator.ID == "" {
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

	w.WriteHeader(202)
}

func (s *Server) CloseAccess(w http.ResponseWriter, r *http.Request) {
	log := logging.GetLogger()

	ctx := r.Context()
	value := ctx.Value(userID)
	userId := value.(string)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var Access models.AccessTo

	Access.UserID = userId

	err = json.Unmarshal(body, &Access)
	if err != nil {
		log.Println(err)
		return
	}

	err = s.Services.CloseAccess(&Access)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(200)
}
