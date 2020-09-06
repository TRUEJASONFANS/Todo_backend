package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Todo struct {
	ID       int64
	Name     string
	Done     bool
	DateTime int64
}

func init() {

}

var instance *TodoDAO

type TodoDAO struct {
	BaseDAO
}

func (dao Todo) Read(todoId int) {
	o := orm.NewOrm()
	user := Todo{ID: int64(todoId)}

	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.ID, user.Name)
	}
}

func (dao TodoDAO) ListAll() []Todo {
	o := orm.NewOrm()
	var todoList []Todo
	num, error := o.Raw("select * from todo").QueryRows(&todoList)
	if error == nil {
		fmt.Println("todo nums: ", num)
	}
	return todoList
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
	temp := *todo
	if o.Read(&temp) == nil {
		if num, err := o.Update(todo); err == nil {
			fmt.Println(num)
			return num
		}
	}
	return -1
}

func (dao TodoDAO) Delete(id int64) int64 {
	o := orm.NewOrm()
	if num, err := o.Delete(&Todo{ID: id}); err == nil {
		return num
	} else {
		fmt.Println(err)
	}
	return -1
}

func GetInstance() *TodoDAO {
	if instance == nil {
		instance = &TodoDAO{}
	}
	return instance
}
