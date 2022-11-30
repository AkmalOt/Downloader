package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type Repository struct {
	Connection *gorm.DB
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{Connection: conn}
}

type Temp struct {
	Name     string
	ID       string
	Password string
}

func (r *Repository) Registration(name string, login string, password []byte) error {
	sqlQwery := `insert into cloud_user(name, login, password)
				values (?,?,?);`

	tx := r.Connection.Exec(sqlQwery, name, login, password)

	if tx.Error != nil {
		log.Println(tx.Error, "help in tx")
		return tx.Error
	}

	return nil
}

func (r *Repository) CreateUserFolder(login string) error {
	var user Temp
	sqlQuery := `select id from cloud_user where login =?;`
	if err := r.Connection.Raw(sqlQuery, login).Scan(&user).Error; err != nil {
		return err
	}

	sqlQwery := `insert into cloud_folders(name, user_id)
				values (?,?);`

	tx := r.Connection.Exec(sqlQwery, login, user.ID)

	if tx.Error != nil {
		log.Println(tx.Error, "help in tx 'CreateUserFolder'")
		return tx.Error
	}

	return nil

}

func (r *Repository) Login(login string) (*Temp, error) {
	var user Temp
	sqlQuery := `select name, password, id from cloud_user where login = ?;`

	if err := r.Connection.Raw(sqlQuery, login).Scan(&user).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	if user.Name == "" {
		log.Println("adfasjkkjafdsdkfjs;fdk;jl")
		return nil, errors.New("Takogo logina net uvazhaemiy")
	}

	//log.Println(user.Name, user.ID, user.Password, "fayt")
	if user.Password == "" {
		return nil, errors.New("Takogo palolya net uvazhaemiy")
	}

	return &user, nil
}

func (r Repository) SetToken(token string, userId string) error {
	//log.Println(token)
	sqlQwery := `insert into cloud_tokens (token, user_id)
				values (?, ?);`

	tx := r.Connection.Exec(sqlQwery, token, userId)
	if tx.Error != nil {
		log.Println("tx error", tx.Error)
		return tx.Error
	}
	return nil
}

type TokenStruct struct {
	ID     string
	UserId string
	Expire time.Time
}

func (r *Repository) ValidateToken(token string) (string, string, error) {
	var tokenS TokenStruct
	log.Println("324")
	sqlQuery := `select *from cloud_tokens where token =?;`

	if err := r.Connection.Raw(sqlQuery, token).Scan(&tokenS).Error; err != nil {
		return "", "", nil
	}

	TimeChecker := time.Now().After(tokenS.Expire)
	log.Println(tokenS.ID, tokenS.Expire)
	if TimeChecker == true {
		return "", "", errors.New(" TimeChecker is true")
	}
	return tokenS.ID, tokenS.UserId, nil
}

//func (r *Repository) UserIdFromToken(token string) (string, error) {
//	var tokenS TokenStruct
//	log.Println("324")
//	sqlQuery := `select *from cloud_tokens where user_id =?;`
//
//	if err := r.Connection.Raw(sqlQuery, token).Scan(&tokenS).Error; err != nil {
//		return "", nil
//	}
//
//	return tokenS.ID, nil
//}

func (r *Repository) FolderCreationForUser(fileName, userId, folderId string) error {

	sqlQwery := `insert into cloud_folders(name, user_id, folder_id)
				values (?, ?, ?);`

	tx := r.Connection.Exec(sqlQwery, fileName, userId, folderId)

	if tx.Error != nil {
		log.Println("tx error", tx.Error)
		return tx.Error
	}
	return nil

}
