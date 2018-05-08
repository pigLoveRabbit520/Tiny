package models

type Article struct {
	Aid          int    `gorm:"primary_key"`
	Title        string `gorm:"type:varchar(200);"`
	Slug         string `gorm:"type:varchar(200);unique_index"`
	Created      uint32 `gorm:"index"`
	Modified     uint32
	Text         string `gorm:"type:longtext"`
	Order        uint
	AuthorId     uint   `gorm:"column:authorId"`
	Template     string `gorm:"type:varchar(32)"`
	Type         string `gorm:"type:varchar(16);default:'post'"`
	Status       string `gorm:"type:varchar(16);default:'publish'"`
	Password     string `gorm:"type:varchar(32)"`
	CommentsNum  uint   `gorm:"column:commentsNum;default:0"`
	AllowComment string `gorm:"type:char(1);column:allowComment;default:0"`
	AllowPing    string `gorm:"type:char(1);column:allowPing;default:0"`
	AllowFeed    string `gorm:"type:char(1);column:allowFeed;default:0"`
	Parent       uint   `gorm:"default:0"`
}
