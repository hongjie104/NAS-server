package models

import (
	"time"

	"github.com/hongjie104/NAS-server/app/pkg/utils"
	"gopkg.in/mgo.v2/bson"
)

// VideoModel VideoModel
type VideoModel struct {
	ID         bson.ObjectId   `bson:"_id,omitempty" json:"id"`
	Code       string          `bson:"code,omitempty" json:"code,omitempty"`
	Name       string          `bson:"name,omitempty" json:"name,omitempty"`
	Date       *time.Time      `bson:"date,omitempty" json:"date,omitempty"`
	Downloaded bool            `bson:"hasDownload,omitempty" json:"hasDownload,omitempty"`
	Score      int             `bson:"score,omitempty" json:"score,omitempty"`
	Actress    []bson.ObjectId `bson:"actress,omitempty" json:"actress,omitempty"`
	Series     bson.ObjectId   `bson:"series,omitempty" json:"series,omitempty"`
	Category   []bson.ObjectId `bson:"category,omitempty" json:"category,omitempty"`
	//     "hasDownload" : false,
}

// VideoIndexOption VideoIndexOption
type VideoIndexOption struct {
	Page      int
	PageSize  int
	Code      string
	ActressID bson.ObjectId
}

// Index 获取影片列表
func (m *VideoModel) Index(option VideoIndexOption) (videoList []VideoModel, total int) {
	ds := NewSessionStore()
	defer ds.Close()

	condition := bson.M{}

	if option.ActressID != "" {
		condition["actress"] = option.ActressID
	}
	if option.Code != "" {
		condition["code"] = bson.M{"$regex": option.Code, "$options": "$i"}
	}

	selector := bson.M{"_id": 1, "name": 1, "code": 1, "date": 1}

	q := ds.C("video").Find(condition).Select(selector)
	total, _ = q.Count()

	if option.Page > 0 && option.PageSize > 0 {
		q = q.Skip((option.Page - 1) * option.PageSize).Limit(option.PageSize)
	}
	q.All(&videoList)
	return
}

// Show a
func (m *VideoModel) Show(id string) (video VideoModel) {
	_id := utils.ToObjectId(id)
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
