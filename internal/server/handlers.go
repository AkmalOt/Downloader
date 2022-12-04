package server

import (
	"Uploader/internal/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
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

	FolderInfo.UserID = userId
	//FolderInfo.FolderID = ""

	err = json.Unmarshal(body, &FolderInfo)
	if err != nil {
		log.Println(err)
		return
	}

	Folders, err := s.Services.GetParentFolders(&FolderInfo)
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

func (s *Server) UploaFile(w http.ResponseWriter, r *http.Request) {

	//==================  PENHUB

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

	log.Println(FileInfo.Name, FileInfo.TargetUrl, FileInfo.UserID, FileInfo.FolderID)

	//
	//bodyBuf := &bytes.Buffer{}

	//bodyWriter := multipart.NewWriter(bodyBuf)

	//fileWriter, err := bodyWriter.CreateFormFile("uploadfile", FileInfo.Name)
	//if err != nil {
	//	fmt.Println("ошибка записи в буфер")
	//	return
	//}

	// процедура открытия файла
	//fh, err := os.Open(FileInfo.Name)
	//if err != nil {
	//	fmt.Println("ошибка открытия файла")
	//	return
	//}

	////iocopy
	//_, err = io.Copy(fileWriter, fh)
	//if err != nil {
	//	return
	//}
	//log.Println("222")
	//contentType := bodyWriter.FormDataContentType()
	//err = bodyWriter.Close()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	////log.Println(FileInfo.TargetUrl, contentType, bodyBuf)
	//resp, err := http.Post(FileInfo.TargetUrl, contentType, bodyBuf)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//
	//log.Println("333")
	//respBody, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//
	//	}
	//}(resp.Body)

	//fmt.Println(resp.Status)
	//fmt.Println(string(respBody))

	err = s.Services.UploadFile(&FileInfo)
	if err != nil {
		log.Println("error in handlers", err)
		return
	}

	//w.WriteHeader(200)

}
