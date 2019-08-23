package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Episode struct {
	ID          int64     `orm:"column(id)"`
	VideoID     string    `orm:"column(video_id)"`
	Title       string    `orm:"column(title)"`
	PlayUrl     string    `orm:"column(play_url)"`
	Source      string    `orm:"column(source)"`
	Description string    `orm:"column(description)"`
	Duration    string    `orm:"column(duration)"`
	CreateTime  time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime  time.Time `orm:"auto_now;type(datetime)"`
}

func (c *Episode) TableName() string {
	return "episode"
}

func init() {
	orm.RegisterModel(new(Episode))
}
