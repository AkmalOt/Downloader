package models

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type AuthInfo struct {
	Name     string `gorm:"name"`
	Login    string `gorm:"login"`
	Password string `gorm:"password"`
}

type Folder struct {
	Name      string `gorm:"name"`
	UserID    string `gorm:"column:user_id"`
	Folder_ID string `gorm:"column:folder_id"`
}
