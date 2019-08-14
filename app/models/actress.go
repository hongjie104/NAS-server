package models

import (
	"time"

	"github.com/hongjie104/NAS-server/app/pkg/utils"

	"gopkg.in/mgo.v2/bson"
)

// ActressModel ActressModel
type ActressModel struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Alias    string        `bson:"alias,omitempty" json:"alias,omitempty"`
	Name     string        `bson:"name" json:"name,omitempty"`
	Birthday *time.Time    `bson:"birthday,omitempty" json:"birthday,omitempty"`
	Score    int           `bson:"score,omitempty" json:"score,omitempty"`
	Height   int           `bson:"height,omitempty" json:"height,omitempty"`
	Bust     int           `bson:"bust,omitempty" json:"bust,omitempty"`
	Waist    int           `bson:"waist,omitempty" json:"waist,omitempty"`
	Hip      int           `bson:"hip,omitempty" json:"hip,omitempty"`
	Cup      string        `bson:"cup,omitempty" json:"cup,omitempty"`
}

// ActressIndexOption ActressIndexOption
type ActressIndexOption struct {
	Page          int
	PageSize      int
	Name          string
	SortBy        string
	ActressIDList []bson.ObjectId
}

// Index 获取女演员列表
func (m *ActressModel) Index(option ActressIndexOption) (actresses []ActressModel, total int) {
	ds := NewSessionStore()
	defer ds.Close()
	condition := bson.M{}
	if option.Name != "" {
		reg := bson.M{"$regex": option.Name, "$options": "$i"}
		condition["$or"] = []bson.M{bson.M{"name": reg}, bson.M{"alias": reg}}
	}
	if option.ActressIDList != nil {
		condition["_id"] = bson.M{"$in": option.ActressIDList}
	}

	sort := ""
	switch option.SortBy {
	case "score-desc":
		sort = "-score"
	}

	selector := bson.M{"_id": 1, "name": 1, "alias": 1, "score": 1}
	q := ds.C("actress").Find(condition).Select(selector)
	total, _ = q.Count()

	if option.Page > 0 && option.PageSize > 0 {
		q = q.Skip((option.Page - 1) * option.PageSize).Limit(option.PageSize)
	}
	if sort != "" {
		q = q.Sort(sort)
	}
	q.All(&actresses)
	return
}

// Show a
func (m *ActressModel) Show(id string) (actress ActressModel) {
	_id := utils.ToObjectId(id)
	ds := NewSessionStore()
	defer ds.Close()
	ds.C("actress").FindId(_id).One(&actress)
	return
}

// Update Update
func (m *ActressModel) Update(id string, data interface{}) {
	_id := bson.ObjectIdHex(id)
	ds := NewSessionStore()
	defer ds.Close()
	ds.C("actress").UpdateId(_id, data)
}
