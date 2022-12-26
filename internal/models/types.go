package models

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type AuthInfo struct {
	Name     string `json:"name" gorm:"name"`
	Login    string `json:"login" gorm:"login"`
	Password string `json:"password" gorm:"password"`
}

type Folder struct {
	ID       string `json:"id" gorm:"id"`
	Name     string `json:"name" gorm:"name"`
	UserID   string `json:"user_id" gorm:"column:user_id"`
	FolderID string `json:"folder_id" gorm:"column:folder_id"`
}

type File struct {
	ID       string `json:"id" gorm:"id"`
	Name     string `json:"name" gorm:"name"`
	UserID   string `json:"user_id" gorm:"column:user_id"`
	FolderID string `json:"folder_id" gorm:"column:folder_id"`
}

type AccessTo struct {
	ID         string `json:"id" gorm:"id"`
	UserID     string `json:"user_id" gorm:"column:user_id"`
	FileId     string `json:"file_id" gorm:"column:file_id"`
	AccessedID string `json:"access_to" gorm:"column:access_to"`
	Active     bool   `json:"active" gorm:"column:active"`
	Expire     string `json:"expire" gorm:"column:expire"`
}

type HumoDataBase struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}
