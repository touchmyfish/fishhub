package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type CrawlerDataMapping struct {
	ID         int64     `orm:"column(id)"`
	VideoID    string    `orm:"column(video_id)"`
	SourceID   string    `orm:"column(source_id)"`
	URI        string    `orm:"column(uri)"`
	Source     string    `orm:"column(source)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

func (c *CrawlerDataMapping) TableName() string {
	return "crawler_data_mapping"
}

func init() {
	orm.RegisterModel(new(CrawlerDataMapping))
}
