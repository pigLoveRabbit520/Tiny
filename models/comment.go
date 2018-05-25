package models

import (
	"fmt"
	"github.com/salamander-mh/SalamanderBlog/config"
	"github.com/salamander-mh/SalamanderBlog/db"
)

type Comment struct {
	Coid     uint   `gorm:"primary_key" json:"coid"`
	Cid      uint   `gorm:"type:int(10) unsigned" json:"cid"`
	Created  uint   `gorm:"type:int(10) unsigned" json:"created"`
	Author   uint   `gorm:"type:varchar(200)" json:"author"`
	AuthorId uint   `gorm:"column:authorId;type:int(10) unsigned" json:"authorId"`
	Mail     string `gorm:"type:varchar(200)" json:"mail"`
	URL      string `gorm:"type:varchar(200)" json:"url"`
	IP       string `gorm:"type:varchar(64)" json:"ip"`
	Agent    string `gorm:"type:varchar(200)" json:"agent"`
	Text     string `gorm:"type:text" json:"text"`
	Type     string `gorm:"type:varchar(16)" json:"type"`
	Status   string `gorm:"type:varchar(16)" json:"logged"`
	Parent   uint   `gorm:"type:int(10) unsigned" json:"parent"`
}

type CommentWithContent struct {
	Comment
	Title string `gorm:"type:varchar(200);" json:"title"`
}

func GetComments(page int) ([]CommentWithContent, int, error) {
	var comments []CommentWithContent
	var total int
	prefix := config.Conf.DB.Prefix
	tbContent := prefix + "contents"
	tbComment := prefix + "comments"
	q := db.DB.Table(fmt.Sprintf("%s com", tbComment)).Select(`com.created, author, mail, url, text, ip, status, 
			parent, con.title, con.cid`).
		Joins("LEFT JOIN "+tbContent+" con ON con.cid = com.cid").
		Where("type = ?", "comment").Order("com.created DESC").Limit(PAGE_SIZE).Offset((page - 1) * PAGE_SIZE).
		Find(&comments)
	if q.Error != nil {
		return nil, -1, q.Error
	}
	db.DB.Table(tbComment).Count(&total)
	return comments, total, nil
}
