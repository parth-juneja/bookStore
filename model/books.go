package model

type Books struct {
	BookId   int    `gorm:"type:int;primary_key"`
	BookName string `gorm:"type:varchar(255)"`
	Author   string `gorm:"type:varchar(255)"`
}
