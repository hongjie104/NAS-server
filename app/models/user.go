package models

import "gopkg.in/mgo.v2/bson"

// UserModel UserModel
type UserModel struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name     string        `bson:"name" json:"name,omitempty"`
	Password string        `bson:"password" json:"password,omitempty"`
}

// Login Login
func (m *UserModel) Login(name string) (u UserModel) {
	ds := NewSessionStore()
	defer ds.Close()
	ds.C("admins").Find(bson.M{"name": name}).One(&u)
	return
}

// Show Show
func (m *UserModel) Show(id string) (u UserModel) {
	ds := NewSessionStore()
	defer ds.Close()
	ds.C("admins").FindId(bson.ObjectIdHex(id)).One(&u)
	return
}
