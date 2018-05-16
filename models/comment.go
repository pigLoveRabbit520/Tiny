package models

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
