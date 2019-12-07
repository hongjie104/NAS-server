package utils

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// SubCn 将字符串 s 从 sub 子串开始截取 len 个字节，支持中文
// 1. 获取子串位置
// 2. 用 []byte 去掉所有子串前面的的字符
// 3. 将剩余部分转化成 rune ，截取 len 个字节
// 4. 最后转化成 string
func SubCn(s, sub string, len int) (string, bool) {
	subIndex := strings.Index(s, sub)
	if subIndex > 0 {
		sByte := []byte(s)[subIndex:]
		sRune := []rune(string(sByte))[0:len]
		return string(sRune), true
	}
	return "", false
}

// ToObjectID ToObjectID
func ToObjectID(val string) (id bson.ObjectId, err error) {
	if bson.IsObjectIdHex(val) {
		id = bson.ObjectIdHex(val)
	} else {
		err = errors.New("id is illegal ObjectId")
	}
	return
}

// Struct2Bson Struct2Bson
func Struct2Bson(in interface{}) (out interface{}, err error) {
	var data []byte
	data, err = bson.Marshal(in)
	if err == nil {
		out = &bson.M{}
		err = bson.Unmarshal(data, out)
	}
	return
}

// RandomInt 随机(0,n]的正整数
func RandomInt(n int) int {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(n)
}

// // CompareFunc CompareFunc
// type CompareFunc func(interface{}, interface{}) int

// // IndexOf IndexOf
// func IndexOf(a []interface{}, e interface{}, cmp CompareFunc) int {
// 	n := len(a)
// 	var i int = 0
// 	for ; i < n; i++ {
// 		if cmp(e, a[i]) == 0 {
// 			return i
// 		}
// 	}
// 	return -1
// }
