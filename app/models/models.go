package models

import (
	"fmt"

	"github.com/hongjie104/NAS-server/app/pkg/setting"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

// Model a
type Model struct {
	CreatedAt  int `bson:"createAt"`
	ModifiedAt int `bson:"modifiedAt"`
}

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
	return s.session.DB(setting.Config.Database.DB).C(name)
}

func init() {
	var err error
	session, err = mgo.Dial(setting.Config.Database.HOST)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect success")
	}

	session.SetMode(mgo.Monotonic, true)
	// session.SetMode(mgo.Eventual, true)
}

// NewSessionStore 为每一HTTP请求创建新的DataStore对象
func NewSessionStore() *SessionStore {
	ds := &SessionStore{
		session.Copy(),
	}
	return ds
}
