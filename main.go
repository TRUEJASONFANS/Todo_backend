package main

import (
	"github.com/astaxie/beego"
	"todo_beego/controllers"
)

func init() {

}

func main() {
	beego.Router("/", &controllers.MainController{})

	//Task
	beego.Router("/task/", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	beego.Router("/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask;"+
		"delete:DeleteTask")

	beego.Run()
}
