package services

import (
	"Uploader/config"
	models "Uploader/internal/models"
	"Uploader/internal/repository"
	logging "Uploader/pkg"
	"github.com/golang-jwt/jwt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Services struct {
	Repository  *repository.Repository
	FileDirPath string
}

func NewServices(rep *repository.Repository) *Services {
	return &Services{Repository: rep, FileDirPath: config.Direction()}
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

// ================================================

var mySingingKey = config.MySingingKey

func (s *Services) Login(userInfo *models.AuthInfo) (string, error) {
	log := logging.GetLogger()

	user, err := s.Repository.Login(userInfo.Login)

	if err != nil {
		log.Println("(s *Services) Login - error", err)
		return "", err
	}

	// ====== Тут идет хеширование

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
	//Тут мы создаем Токен.

	// старый метод
	//buf := make([]byte, 256)

	//_, err = rand.Read(buf)
	//if err != nil {
	//	return "", err
	//}

	//token := hex.EncodeToString(buf)

	// JWT Method
	token := jwt.New(jwt.SigningMethodHS384)

	claims := token.Claims.(jwt.MapClaims)
	claims["auth"] = true
	claims["user"] = "test"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	log.Println(string(mySingingKey))
	tokenString, err := token.SignedString(mySingingKey)
	if err != nil {
		log.Println(err)
		return "", err
	}

	err = s.Repository.SetToken(tokenString, user.ID)
	if err != nil {
		log.Println(err)
		return "", nil
	}
	return tokenString, nil

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
	//log.Println(files)
	return files, nil
}

func (s *Services) SaveFile(file io.Reader, fileName string, upfile *models.File) (*models.File, error) {
	log := logging.GetLogger()
	log.Print(fileName)
	var dot int

	switch {
	case string(fileName[len(fileName)-5]) == ".":
		dot = 5
	case string(fileName[len(fileName)-4]) == ".":
		dot = 4
	}

	extension := fileName[len(fileName)-dot:]
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

//func (s *Services) ValidationForDownload(files *models.File) (string, error) { // todo need to rework
//	ab, err := s.ValidationForDownload(files)
//	if err != nil {
//		log.Println(err)
//		return "", err
//	}
//	log.Println(ab)
//	return ab, nil
//}

func (s *Services) ChangeFileName(files *models.File) error {
	return s.Repository.ChangeFileName(files)
}

func (s *Services) ChangeFolderName(folder *models.Folder) error {
	return s.Repository.ChangeFolderName(folder)
}

func (s *Services) GetFileInfoByID(files *models.File) (*models.File, error) {
	return s.Repository.GetFileInfoByID(files)
}

func (s *Services) DeleteFile(files *models.File) error {
	return s.Repository.DeleteFile(files)
}

func (s *Services) GiveAccess(file *models.AccessTo) error {
	return s.Repository.GiveAccess(file)
}

func (s *Services) GetAccessedFiles(file *models.AccessTo) ([]*models.AccessTo, error) {
	return s.Repository.GetAccessedFiles(file)
}

func (s *Services) DownloadAccessedFiles(file *models.File) (*models.File, error) {
	return s.Repository.DownloadAccessedFiles(file.ID)
}

func (s *Services) ValidationForAccessDownload(file *models.AccessTo) (*models.AccessTo, error) {
	return s.Repository.ValidationForAccessDownload(file)
}

func (s *Services) CloseAccess(file *models.AccessTo) error {
	return s.Repository.CloseAccess(file)
}
