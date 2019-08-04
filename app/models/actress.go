package models

import "time"

// Actress 女演员
type Actress struct {
	// Model

	// Name string `bson:"name"`
	// Age  int    `bson:"age"`
	ID       string    `bson:"_id"`
	Alias    string    `bson:"alias"`
	Birthday time.Time `bson:"birthday"`

	// 	{

	//     "height" : 0,
	//     "bust" : 0,
	//     "waist" : 0,
	//     "hip" : 0,
	//     "cup" : "X",
	//     "score" : 0,
	//     "javBusCode" : "12f",
	//     "name" : "田嶋涼子",
	//     "javBusNum" : 1383,
	//     "dmmScore" : 0,
	//     "video" : [
	//         ObjectId("5a96a69e77480a048a31e0d5"),
	//         ObjectId("5a96a6f477480a048a31e0ee"),
	//         ObjectId("5a9be92477480a3114e4e891"),
	//         ObjectId("5a9bf66377480a3114e4ec86"),
	//         ObjectId("5a9bf66877480a3114e4ec88"),
	//         ObjectId("5a9bf66b77480a3114e4ec8a"),
	//         ObjectId("5a9bf66e77480a3114e4ec8b"),
	//         ObjectId("5a9bf67177480a3114e4ec8c")
	//     ],
	//     "t2s" : true,
	//     "numVideo" : 8
	// }
}

// Index 获取女演员列表
func Index(page int, pageSize int) (actress []Actress) {
	db := NewSessionStore()
	defer db.Close()
	db.C("actress").Find(nil).Skip((page - 1) * pageSize).Limit(pageSize).All(&actress)
	return
}
