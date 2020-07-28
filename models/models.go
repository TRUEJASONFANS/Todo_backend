package models

import (
	"fmt"
)

type User struct {
	Id          int
	Name        string
	Profile     *Profile `orm:"rel(one)"`      // OneToOne relation
	Post        []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id          int
	Age         int16
	User        *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"` //设置多对多反向关系
}

func init() {
	// 需要在init中注册定义的model
	fmt.Println("init the User models")
	//orm.RegisterModel(new(User), new(Post), new(Profile), new(Tag))
}
