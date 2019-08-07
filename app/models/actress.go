package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// ActressModel ActressModel
type ActressModel struct{}

// ActressListData 女演员
type ActressListData struct {
	ID       bson.ObjectId `bson:"_id"`
	Alias    string        `bson:"alias"`
	Birthday time.Time     `bson:"birthday"`
}

// ActressData 女演员
type ActressData struct {
	ActressListData
	Name string `bson:"name"`
}

// Index 获取女演员列表
func (m ActressModel) Index(page int, pageSize int) (actresses []ActressListData) {
	ds := NewSessionStore()
	defer ds.Close()
	ds.C("actress").Find(nil).Select(bson.M{"_id": 1, "alias": 1}).Skip((page - 1) * pageSize).Limit(pageSize).All(&actresses)
	return
}

// Show a
func (m ActressModel) Show(id string) (actress ActressData) {
	if bson.IsObjectIdHex(id) {
		_id := bson.ObjectIdHex(id)
		ds := NewSessionStore()
		defer ds.Close()
		ds.C("actress").FindId(_id).One(&actress)
		// ds.C("actress").Find(bson.M{"_id": _id}).One(&actress)
		return
	}
	return
}
