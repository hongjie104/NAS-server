package model

import (
	"time"

	"github.com/hongjie104/NAS-server/pkg/utils"
	"gopkg.in/mgo.v2/bson"
)

// VideoModel VideoModel
type VideoModel struct {
	ID         bson.ObjectId   `bson:"_id,omitempty" json:"id"`
	Code       string          `bson:"code,omitempty" json:"code,omitempty"`
	Name       string          `bson:"name,omitempty" json:"name,omitempty"`
	Date       *time.Time      `bson:"date,omitempty" json:"date,omitempty"`
	Subtitle   bool            `bson:"subtitle,omitempty" json:"subtitle,omitempty"` // 有没有字幕
	HD         bool            `bson:"hd,omitempty" json:"hd,omitempty"`             // 是不是高清的
	Score      int             `bson:"score,omitempty" json:"score,omitempty"`
	Actress    []bson.ObjectId `bson:"actress,omitempty" json:"actress,omitempty"`
	Series     bson.ObjectId   `bson:"series,omitempty" json:"series,omitempty"`
	Category   []bson.ObjectId `bson:"category,omitempty" json:"category,omitempty"`
	Image      string          `bson:"img,omitempty" json:"img,omitempty"`
	ImageCover string          `bson:"img_s,omitempty" json:"img_s,omitempty"`
}

// VideoIndexOption VideoIndexOption
type VideoIndexOption struct {
	Page      int
	PageSize  int
	Code      string
	ActressID string
	SeriesID  string
}

// Index 获取影片列表
func (m *VideoModel) Index(option VideoIndexOption) (videoList []VideoModel, total int) {
	ds := NewSessionStore()
	defer ds.Close()

	condition := bson.M{}

	if option.ActressID != "" {
		condition["actress"], _ = utils.ToObjectID(option.ActressID)
	}
	if option.Code != "" {
		condition["code"] = bson.M{"$regex": option.Code, "$options": "$i"}
	}
	if option.SeriesID != "" {
		condition["series"], _ = utils.ToObjectID(option.SeriesID)
	}

	selector := bson.M{"_id": 1, "name": 1, "code": 1, "date": 1, "img_s": 1, "subtitle": 1, "hd": 1}

	q := ds.C("video").Find(condition).Select(selector)
	total, _ = q.Count()

	if option.Page > 0 && option.PageSize > 0 {
		q = q.Skip((option.Page - 1) * option.PageSize).Limit(option.PageSize)
	}
	q.Sort("-date").All(&videoList)
	return
}

// Show a
func (m *VideoModel) Show(id string) (video VideoModel) {
	_id, _ := utils.ToObjectID(id)
	ds := NewSessionStore()
	defer ds.Close()
	ds.C("video").FindId(_id).One(&video)
	return
}

// Update Update
func (m *VideoModel) Update(id string, data interface{}) {
	_id := bson.ObjectIdHex(id)
	ds := NewSessionStore()
	defer ds.Close()
	ds.C("video").UpdateId(_id, data)
}
