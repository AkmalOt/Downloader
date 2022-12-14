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

//type FileLinkInSrv struct {
//	FileDirPath string
//}

type FilesAndFolders struct {
	Folder
	File
}
