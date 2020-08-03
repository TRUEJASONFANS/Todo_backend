package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"

	TaskDao "todo_beego/dao"
)

var DefaultTaskList *TaskManager
var dao = TaskDao.GetInstance()

// NewTask creates a new task given a title, that can't be empty.
func NewTask(title string) (*TaskDao.Todo, error) {
	if title == "" {
		return nil, fmt.Errorf("empty title")
	}
	return &TaskDao.Todo{ID: 0, Name: title, Done: false, DateTime: time.Now().Unix()}, nil
}

// TaskManager manages a list of tasks in memory.
type TaskManager struct {
	tasks  []*TaskDao.Todo
	lastID int64
}

// NewTaskManager returns an empty TaskManager.
func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

// Save saves the given Task in the TaskManager.
func (m *TaskManager) Save(task *TaskDao.Todo) error {

	if task.ID == 0 {
		m.lastID++
		task.ID = m.lastID
		var clonedTask = cloneTask(task)
		m.tasks = append(m.tasks, clonedTask)
		dao.Create(clonedTask)
		return nil
	}

	for i, t := range m.tasks {
		if t.ID == task.ID {
			m.tasks[i] = cloneTask(task)
			dao.Update(m.tasks[i])
			return nil
		}
	}
	return fmt.Errorf("unknown task")
}

// cloneTask creates and returns a deep copy of the given Task.
func cloneTask(t *TaskDao.Todo) *TaskDao.Todo {
	c := *t
	return &c
}

// All returns the list of all the Tasks in the TaskManager.
func (m *TaskManager) All() []*TaskDao.Todo {
	return m.tasks
}

// Find returns the Task with the given id in the TaskManager and a boolean
// indicating if the id was found.
func (m *TaskManager) Find(ID int64) (*TaskDao.Todo, bool) {
	for _, t := range m.tasks {
		if t.ID == ID {
			return t, true
		}
	}
	return nil, false
}

func (m *TaskManager) Delete(id int64) error {
	var idToRemoved = -1
	for index, t := range m.tasks {
		if t.ID == id {
			idToRemoved = index
			dao.Delete(id)
		}
	}
	if idToRemoved != -1 {
		m.tasks = append(m.tasks[:idToRemoved], m.tasks[idToRemoved+1:]...)
	}
	return nil

}

func init() {
	DefaultTaskList = NewTaskManager()
	m := DefaultTaskList
	todoList := TaskDao.GetInstance().ListAll()
	m.tasks = make([]*TaskDao.Todo, 0)
	for i, _ := range todoList {
		m.tasks = append(m.tasks, &todoList[i])
	}
}
