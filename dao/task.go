package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Todo struct {
	Id          int
	Name        string
	Done  bool
	dateTime  int64
}



func init() {

}

var instance *TodoDAO

type TodoDAO struct {
	BaseDAO
}

func (dao Todo) Read(todoId int) {
	o := orm.NewOrm()
	user := Todo{Id: todoId}

	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Id, user.Name)
	}
}

func (dao TodoDAO) Create(todo *Todo) {
	o := orm.NewOrm()
	id, err := o.Insert(todo)
	if err == nil {
		fmt.Println(id)
	}
}

func (dao TodoDAO) Update(todo *Todo) int64 {
	o := orm.NewOrm()
	if o.Read(&todo) == nil {
		if num, err := o.Update(&todo); err == nil {
			fmt.Println(num)
			return num
		}
	}
	return -1
}

func (dao TodoDAO) Delete(todo *Todo) int64 {
	o := orm.NewOrm()
	if num, err := o.Delete(&todo); err == nil {
		return num
	} else {
		fmt.Println(err)
	}
	return -1
}

func GetInstance() *TodoDAO {
	if instance == nil {
		instance = &TodoDAO {}
	}
	return instance
}
