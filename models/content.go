package models

import (
	"fmt"
	"github.com/salamander-mh/SalamanderBlog/config"
	"github.com/salamander-mh/SalamanderBlog/db"
)

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

const (
	PAGE_SIZE = 15
)

func getLimit(page int) string {
	return fmt.Sprintf("LIMIT %d, %d", (page-1)*PAGE_SIZE, PAGE_SIZE)
}

type ContentForAdmin struct {
	Cid         int    `json:"cid"`
	Title       string `json:"title"`
	Created     uint32 `json:"created"`
	ScreenName  string `json:"screenName" gorm:"column:screenName"`
	CommentsNum uint   `json:"commentsNum" gorm:"column:commentsNum"`
	MetaCat     string `json:"meta_cat" gorm:"column:meta_cat"`
}

func GetPosts(page int) ([]ContentForAdmin, int, error) {
	prefix := config.Conf.DB.Prefix
	tbContent := prefix + "contents"
	tbUser := prefix + "users"
	tbRelation := prefix + "relationships"
	var posts []ContentForAdmin
	var postsCount int
	q := db.DB.Raw(fmt.Sprintf(`SELECT cid, title, c.created, commentsNum, u.screenName,
		(SELECT GROUP_CONCAT(mid) FROM %s r WHERE r.cid = c.cid) AS meta_cat		
		FROM %s c LEFT JOIN %s u ON c.authorId = u.uid WHERE c.type = 'post' ORDER BY created DESC %s`,
		tbRelation, tbContent, tbUser, getLimit(page))).
		Find(&posts)
	if q.Error != nil {
		return nil, -1, q.Error
	}
	db.DB.Table(tbContent).Where("type = ?", "post").Count(&postsCount)
	return posts, postsCount, nil
}

type Page struct {
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	ScreenName string `json:"screenName" gorm:"column:screenName"`
	Created    uint32 `json:"created"`
}

func GetPages(page int) ([]Page, int, error) {
	var pages []Page
	var pagesCount int
	tbUser := config.Conf.DB.Prefix + "users"
	tbContent := config.Conf.DB.Prefix + "contents"
	q := db.DB.Table(tbContent+" c").Select("title, slug, c.created, u.screenName").Joins(fmt.Sprintf("LEFT JOIN %s u ON u.uid = c.authorId", tbUser)).
		Where("type = ?", "page").
		Limit(PAGE_SIZE).Offset((page - 1) * PAGE_SIZE).Find(&pages)
	if q.Error != nil {
		return nil, -1, q.Error
	}
	q.Offset(-1).Limit(-1).Count(&pagesCount)
	return pages, pagesCount, nil
}
