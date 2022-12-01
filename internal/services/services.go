package services

import (
	"Uploader/internal/models"
	"Uploader/internal/repository"
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

type Services struct {
	Repository *repository.Repository
}

func NewServices(rep *repository.Repository) *Services {
	return &Services{Repository: rep}
}

// Регистарция
func (s *Services) Register(userInfo *models.AuthInfo) error {
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
	log.Println("test in ShowFolder of service", folder[1])
	return list, err
}

func (s *Services) GetParentFolders(userInfo *models.Folder) ([]*models.Folder, error) {
	var list []*models.Folder
	folder, err := s.Repository.GetParentFolders(userInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	list = folder
	log.Println("test in ShowFolder of service", folder[1])
	return list, err
}
