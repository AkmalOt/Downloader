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
	Name     string `json:"name" gorm:"name"`
	UserID   string `json:"user_id" gorm:"column:user_id"`
	FolderID string `json:"folder_id" gorm:"column:folder_id"`
}

type File struct {
	Name      string `json:"name" gorm:"name"`
	TargetUrl string `json:"url" gorm:"url"`
	UserID    string `json:"user_id" gorm:"column:user_id"`
	FolderID  string `json:"folder_id" gorm:"column:folder_id"`
}
