package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
	"todo_beego/controllers"
)

func init() {

}

func main() {

	var success = []byte("SUPPORT OPTIONS")

	var corsFunc = func(ctx *context.Context) {
		origin := ctx.Input.Header("Origin")
		ctx.Output.Header("Access-Control-Allow-Methods", "OPTIONS,DELETE,POST,GET,PUT,PATCH")
		ctx.Output.Header("Access-Control-Max-Age", "3600")
		ctx.Output.Header("Access-Control-Allow-Headers", "X-Custom-Header,accept,Content-Type,Access-Token")
		ctx.Output.Header("Access-Control-Allow-Credentials", "true")
		ctx.Output.Header("Access-Control-Allow-Origin", origin)
		if ctx.Input.Method() == http.MethodOptions {
			// options请求，返回200
			ctx.Output.SetStatus(http.StatusOK)
			_ = ctx.Output.Body(success)
		}
	}

	beego.InsertFilter("/*", beego.BeforeRouter, corsFunc)

	beego.Router("/", &controllers.MainController{})

	//Task

	beego.Router("/todos/api/task", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	beego.Router("/todos/api/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask;"+
		"delete:DeleteTask")

	beego.Run()
}
