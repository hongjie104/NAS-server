package model

import (
	"fmt"
	"sync"

	"github.com/hongjie104/NAS-server/config"
	"gopkg.in/mgo.v2"
)

var once sync.Once

var session *mgo.Session

// SessionStore a
type SessionStore struct {
	session *mgo.Session
}

// Close a
func (s *SessionStore) Close() {
	s.session.Close()
}

// C a
func (s *SessionStore) C(name string) *mgo.Collection {
	return s.session.DB(config.Config.Database.DB).C(name)
}

// UserModelInstance UserModelInstance
var UserModelInstance *UserModel

// ActressModelInstance ActressModelInstance
var ActressModelInstance *ActressModel

// VideoModelInstance VideoModelInstance
var VideoModelInstance *VideoModel

// SeriesModelInstance SeriesModelInstance
var SeriesModelInstance *SeriesModel

// CategoryModelInstance CategoryModelInstance
var CategoryModelInstance *CategoryModel

func init() {
	var err error
	session, err = mgo.Dial(config.Config.Database.HOST)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect success")
	}

	session.SetMode(mgo.Monotonic, true)
	// session.SetMode(mgo.Eventual, true)

	UserModelInstance = &UserModel{}
	ActressModelInstance = &ActressModel{}
	VideoModelInstance = &VideoModel{}
	SeriesModelInstance = &SeriesModel{}
	CategoryModelInstance = &CategoryModel{}
}

// NewSessionStore 为每一HTTP请求创建新的DataStore对象
func NewSessionStore() *SessionStore {
	ds := &SessionStore{
		session.Copy(),
	}
	return ds
}
