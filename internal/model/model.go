package model

type Model struct {
	ID    int64 `gorm:"primary_key"  json:"id"`
	Ctime int64 `json:"ctime"`
	Mtime int64 `json:"mtime"`
	IsDel int64 `json:"is_del"`
}
