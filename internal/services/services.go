package services

import (
	models "Uploader/internal/models"
	"Uploader/internal/repository"
	logging "Uploader/pkg"
	"crypto/rand"
	"encoding/hex"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Services struct {
	Repository  *repository.Repository
	FileDirPath string //todo change, add to configs
}

func NewServices(rep *repository.Repository) *Services {
	return &Services{Repository: rep, FileDirPath: "./files"}
}

// Регистарция
func (s *Services) Register(userInfo *models.AuthInfo) error {
	log := logging.GetLogger()
	hash, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = s.Repository.Registration(userInfo.Name, userInfo.Login, hash)
	if err != nil {
		return err
	}

	err = s.Repository.CreateUserFolder(userInfo.Login)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *Services) Login(userInfo *models.AuthInfo) (string, error) {
	log := logging.GetLogger()

	user, err := s.Repository.Login(userInfo.Login)

	if err != nil {
		log.Println("(s *Services) Login - error", err)
		return "", err
	}

	//log.Println(user.Password, " ----", userInfo.Password)

	hash, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInfo.Password))
	if err != nil {
		log.Println("error in CompareHash")
		return "", err
	}

	log.Println(user.Password, " ----", userInfo.Password, string(hash))

	//============================================================================
	buf := make([]byte, 256)

	_, err = rand.Read(buf)
	if err != nil {
		return "", err
	}

	token := hex.EncodeToString(buf)

	err = s.Repository.SetToken(token, user.ID)
	if err != nil {
		log.Println(err)
		return "", nil
	}
	return token, nil

}

func (s *Services) FolderCreation(userInfo *models.Folder) error {
	log := logging.GetLogger()
	err := s.Repository.FolderCreationForUser(userInfo.Name, userInfo.UserID, userInfo.FolderID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Services) GetFoldersFromParent(userInfo *models.Folder) ([]*models.Folder, error) {
	var list []*models.Folder
	folder, err := s.Repository.GetFoldersFromParent(userInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	list = folder
	//log.Println("test in ShowFolder of service", folder[1])
	return list, err
}

func (s *Services) GetParentFolders(userInfo *models.Folder) (string, []*models.Folder, error) {
	log := logging.GetLogger()
	//var list []*models.Folder
	id, folder, err := s.Repository.GetParentFolders(userInfo)
	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	//list = folder
	//log.Println("test in ShowFolder of service", folder[1])
	return id, folder, err
}

func (s *Services) GetFiles(userFiles *models.File) ([]*models.File, error) {
	log := logging.GetLogger()
	files, err := s.Repository.GetFiles(userFiles)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(files)
	return files, nil
}

func (s *Services) SaveFile(file io.Reader, fileName string, upfile *models.File) (*models.File, error) {
	log := logging.GetLogger()
	log.Print(fileName)
	extension := fileName[len(fileName)-4:]
	var count int
	for i := 0; i <= len(fileName); i++ {
		count++
		if string(fileName[i]) == "." {

			break
		}
	}
	upfile.Name = fileName[0 : count-1]
	path := filepath.Join(s.FileDirPath, upfile.Name+extension)
	MainFile, err := os.Create(path)
	if err != nil {
		err := errors.WithStack(err)
		return nil, err
	}

	defer MainFile.Close()

	_, err = io.Copy(MainFile, file)
	if err != nil {
		//err := errors.WithStack(err)
		log.Println(err)
		return nil, err
	}

	upfile.Name = upfile.Name + extension
	//log.Println("tetetetetetet", upfile.Name, upfile.UserID, upfile.TargetUrl)
	log.Println(upfile.Name, upfile.UserID)
	return upfile, nil
}

func (s *Services) UploadFile(file *models.File) error {
	log := logging.GetLogger()

	err := s.Repository.UploadFile(file.Name, file.UserID, file.FolderID)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (s *Services) DownloadFiles(id string) (*models.File, error) {
	
	return s.Repository.DownloadFiles(id)
}

//func (s *Services) ValidationForDownload(file *models.File) (string, error) { // todo need to rework
//	ab, err := s.ValidationForDownload(file)
//	if err != nil {
//		log.Println(err)
//		return "", err
//	}
//	log.Println(ab)
//	return ab, nil
//}
