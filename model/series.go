package model

import (
	"gopkg.in/mgo.v2/bson"
)

// SeriesModel SeriesModel
type SeriesModel struct {
	ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string        `bson:"name" json:"name,omitempty"`
}

// Index 获取系列列表
func (m *SeriesModel) Index(page, pageSize int, name string) (series []SeriesModel, total int) {
	ds := NewSessionStore()
	defer ds.Close()
	condition := bson.M{}
	if name != "" {
		condition["name"] = bson.M{"$regex": name, "$options": "$i"}
	}

	q := ds.C("series").Find(condition)
	total, _ = q.Count()

	if page > 0 && pageSize > 0 {
		q = q.Skip((page - 1) * pageSize).Limit(pageSize)
	}
	q.All(&series)
	return
}

// Show a
func (m *SeriesModel) Show(id bson.ObjectId) (series SeriesModel) {
	ds := NewSessionStore()
	defer ds.Close()
	ds.C("series").FindId(id).One(&series)
	return
}

// // Update Update
// func (m *SeriesModel) Update(id string, data interface{}) {
// 	_id := bson.ObjectIdHex(id)
// 	ds := NewSessionStore()
// 	defer ds.Close()
// 	ds.C("actress").UpdateId(_id, data)
// }
