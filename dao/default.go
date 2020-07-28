package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:1234@tcp(127.0.0.1:3306)/todos_db?charset=utf8")
	fmt.Println("loading the database driver")
	fmt.Println("init the task dao")
	orm.RegisterModel(new(Todo))

	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	// 数据库别名
	name := "default"

	force := true

	// 打印执行过程
	verbose := true

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

type BaseDAO struct {

}