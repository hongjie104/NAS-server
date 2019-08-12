package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// VideoModel VideoModel
type VideoModel struct {
	ID         bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Code       string        `bson:"code" json:"code"`
	Name       string        `bson:"name" json:"name"`
	Date       *time.Time    `bson:"date" json:"date"`
	Downloaded bool          `bson:"hasDownload" json:"hasDownload"`
	Score      int           `bson:"score" json:"score"`

	// 	{
	//     "category" : [
	//         ObjectId("e598880130ba5c943a060000"),
	//         ObjectId("9ad18c0130ba5cbc358d0100"),
	//         ObjectId("98658c0130ba5cbc351a0000")
	//     ],
	//     "series" : ObjectId("5a9bf7de77480a3114e4ece5"),
	//     "actress" : [
	//         ObjectId("5a913ddb77480a35c0710331"),
	//         ObjectId("5a91390d77480a35c0710139")
	//     ],
	//     "score" : 0,
	//     "hasDownload" : false,
	// }
}

// Index 获取影片列表
func (m *VideoModel) Index(page int, pageSize int) (videoList []VideoModel, total int) {
	ds := NewSessionStore()
	defer ds.Close()
	ds.C("video").Find(nil).Select(bson.M{"_id": 1, "name": 1, "code": 1, "date": 1}).Skip((page - 1) * pageSize).Limit(pageSize).All(&videoList)
	total, _ = ds.C("video").Count()
	return
}

// Show a
func (m *VideoModel) Show(id string) (video VideoModel) {
	// if bson.IsObjectIdHex(id) {
	_id := bson.ObjectIdHex(id)
	ds := NewSessionStore()
	defer ds.Close()
	ds.C("video").FindId(_id).One(&video)
	return
	// }
	// return
}

// Update Update
func (m *VideoModel) Update(id string, data interface{}) {
	_id := bson.ObjectIdHex(id)
	ds := NewSessionStore()
	defer ds.Close()
	ds.C("video").UpdateId(_id, data)
}
