package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// ActressModel ActressModel
type ActressModel struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Alias    string        `bson:"alias,omitempty" json:"alias,omitempty"`
	Name     string        `bson:"name" json:"name,omitempty"`
	Birthday *time.Time    `bson:"birthday,omitempty" json:"birthday,omitempty"`
}

// Index 获取女演员列表
func (m *ActressModel) Index(page int, pageSize int) (actresses []ActressModel) {
	ds := NewSessionStore()
	defer ds.Close()
	ds.C("actress").Find(nil).Select(bson.M{"_id": 1, "name": 1, "alias": 1}).Skip((page - 1) * pageSize).Limit(pageSize).All(&actresses)
	return
}

// Show a
func (m *ActressModel) Show(id string) (actress ActressModel) {
	// if bson.IsObjectIdHex(id) {
	_id := bson.ObjectIdHex(id)
	ds := NewSessionStore()
	defer ds.Close()
	ds.C("actress").FindId(_id).One(&actress)
	return
	// }
	// return
}

// Update Update
func (m *ActressModel) Update(id string, data interface{}) {
	_id := bson.ObjectIdHex(id)
	ds := NewSessionStore()
	defer ds.Close()
	ds.C("actress").UpdateId(_id, data)
}
