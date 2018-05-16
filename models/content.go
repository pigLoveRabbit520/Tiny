package models

type Content struct {
	Cid          int    `gorm:"primary_key" json:"cid"`
	Title        string `gorm:"type:varchar(200);" json:"title"`
	Slug         string `gorm:"type:varchar(200);unique_index" json:"slug"`
	Created      uint32 `gorm:"index" json:"created"`
	Modified     uint32
	Text         string `gorm:"type:longtext" json:"text"`
	Order        uint   `gorm:"type:int(10) unsigned" json:"order"`
	AuthorId     uint   `gorm:"column:authorId;type:int(10) unsigned" json:"authorId"`
	Template     string `gorm:"type:varchar(32)" json:"template"`
	Type         string `gorm:"type:varchar(16);default:'post'" json:"type"`
	Status       string `gorm:"type:varchar(16);default:'publish'" json:"status"`
	Password     string `gorm:"type:varchar(32)"`
	CommentsNum  uint   `gorm:"column:commentsNum;default:0"`
	AllowComment string `gorm:"type:char(1);column:allowComment;default:0"`
	AllowPing    string `gorm:"type:char(1);column:allowPing;default:0"`
	AllowFeed    string `gorm:"type:char(1);column:allowFeed;default:0"`
	Parent       uint   `gorm:"default:0"`
}
