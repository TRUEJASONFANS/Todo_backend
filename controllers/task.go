package controllers

import (
	"encoding/json"
	"strconv"
	TaskDao "todo_beego/dao"
	"todo_beego/models"

	"github.com/astaxie/beego"
)

type TaskController struct {
	beego.Controller
}

// Example:
//
//   req: GET /task/
//   res: 200 {"Tasks": [
//          {"ID": 1, "Title": "Learn Go", "Done": false},
//          {"ID": 2, "Title": "Buy bread", "Done": true}
//        ]}
func (task *TaskController) ListTasks() {
	res := struct{ Tasks []*TaskDao.Todo }{models.DefaultTaskList.All()}
	task.Data["json"] = res
	task.ServeJSON()
}

// Examples:
//
//   req: POST /task/ {"Title": ""}
//   res: 400 empty title
//
//   req: POST /task/ {"Title": "Buy bread"}
//   res: 200
func (task *TaskController) NewTask() {
	req := struct{ Name string }{}
	if err := json.Unmarshal(task.Ctx.Input.RequestBody, &req); err != nil {
		task.Ctx.Output.SetStatus(400)
		task.Ctx.Output.Body([]byte("empty Name"))
		return
	}
	t, err := models.NewTask(req.Name)
	if err != nil {
		task.Ctx.Output.SetStatus(400)
		task.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	models.DefaultTaskList.Save(t)
	task.Data["json"] = &t
	task.ServeJSON()
}

// Examples:
//
//   req: GET /task/1
//   res: 200 {"ID": 1, "Title": "Buy bread", "Done": true}
//
//   req: GET /task/42
//   res: 404 task not found
func (task *TaskController) GetTask() {
	id := task.Ctx.Input.Param(":id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	t, ok := models.DefaultTaskList.Find(intid)
	beego.Info("Found", ok)
	if !ok {
		task.Ctx.Output.SetStatus(404)
		task.Ctx.Output.Body([]byte("task not found"))
		return
	}
	task.Data["json"] = t
	task.ServeJSON()
}

// Example:
//
//   req: PUT /task/1 {"ID": 1, "Title": "Learn Go", "Done": true}
//   res: 200
//
//   req: PUT /task/2 {"ID": 2, "Title": "Learn Go", "Done": true}
//   res: 400 inconsistent task IDs
func (task *TaskController) UpdateTask() {
	id := task.Ctx.Input.Param(":id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	var t TaskDao.Todo
	if err := json.Unmarshal(task.Ctx.Input.RequestBody, &t); err != nil {
		task.Ctx.Output.SetStatus(400)
		task.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	if t.ID != intid {
		task.Ctx.Output.SetStatus(400)
		task.Ctx.Output.Body([]byte("inconsistent task IDs"))
		return
	}
	if _, ok := models.DefaultTaskList.Find(intid); !ok {
		task.Ctx.Output.SetStatus(400)
		task.Ctx.Output.Body([]byte("task not found"))
		return
	}
	models.DefaultTaskList.Save(&t)
	res := struct{ task TaskDao.Todo }{t}
	task.Data["json"] = res
	task.ServeJSON()
}

func (task *TaskController) DeleteTask() {
	id := task.Ctx.Input.Param(":id")
	beego.Info("Task is ", id)
	intid, _ := strconv.ParseInt(id, 10, 64)
	if _, ok := models.DefaultTaskList.Find(intid); !ok {
		task.Ctx.Output.SetStatus(400)
		task.Ctx.Output.Body([]byte("task not found"))
		return
	}
	models.DefaultTaskList.Delete(intid)
}
