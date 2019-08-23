package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Video struct {
	ID            int64     `orm:"column(id)"`
	VideoID       string    `orm:"column(video_id)"`
	Title         string    `orm:"column(title)"`
	Series        string    `orm:"column(series)"`
	Area          string    `orm:"column(area)"`
	Status        string    `orm:"column(status)"`
	Captions      string    `orm:"column(captions)"`
	Director      string    `orm:"column(director)"`
	Actor         string    `orm:"column(actor)"`
	PlotCategory  string    `orm:"column(plot_category)"`
	LevelCategory string    `orm:"column(level_category)"`
	OtherCategory string    `orm:"column(other_category)"`
	Description   string    `orm:"column(description)"`
	PublishedTime time.Time `orm:"column(published_time)"`
	CreateTime    time.Time `orm:"auto_now_add;type(create_time)"`
	UpdateTime    time.Time `orm:"auto_now;type(update_time)"`
}

func (c *Video) TableName() string {
	return "video"
}

func init() {
	orm.RegisterModel(new(Video))
}
