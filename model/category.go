package model

import (
	"gopkg.in/mgo.v2/bson"
)

// CategoryModel CategoryModel
type CategoryModel struct {
	ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string        `bson:"name" json:"name,omitempty"`
}

// Index 获取女演员列表
func (m *CategoryModel) Index(idArr []bson.ObjectId) (category []CategoryModel, total int) {
	if len(idArr) < 1 {
		return
	}
	ds := NewSessionStore()
	defer ds.Close()
	condition := bson.M{"_id": bson.M{"$in": idArr}}

	q := ds.C("category").Find(condition)
	total, _ = q.Count()

	q.All(&category)
	return
}

// // Show a
// func (m *CategoryModel) Show(id bson.ObjectId) (category CategoryModel) {
// 	ds := NewSessionStore()
// 	defer ds.Close()
// 	ds.C("category").FindId(id).One(&category)
// 	return
// }

// ShowMany ShowMany
func (m *CategoryModel) ShowMany(idList []bson.ObjectId) (categoryList []CategoryModel) {
	if idList == nil || len(idList) < 1 {
		return
	}
	ds := NewSessionStore()
	defer ds.Close()
	ds.C("category").Find(bson.M{"_id": bson.M{"$in": idList}}).All(&categoryList)
	return
}

// // Update Update
// func (m *CategoryModel) Update(id string, data interface{}) {
// 	_id := bson.ObjectIdHex(id)
// 	ds := NewSessionStore()
// 	defer ds.Close()
// 	ds.C("actress").UpdateId(_id, data)
// }
