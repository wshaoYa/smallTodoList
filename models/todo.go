package models

import (
	"smallTodoList/dao"
)

type Todo struct {
	Id     int    `json:"id"`
	Task   string `json:"title"`
	Status bool   `json:"status"`
}

// CreateTodo 新建一个Todo
func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	return
}

// DeleteTodo 删除一个Todo
func DeleteTodo(id string) (err error) {
	err = dao.DB.Delete(&Todo{}, id).Error
	return
}

// UpdateTodoStatus 修改Todo状态
func UpdateTodoStatus(id string) (err error) {
	var todo Todo
	dao.DB.First(&todo, id)
	todo.Status = !todo.Status
	err = dao.DB.Save(&todo).Error
	return
}

// GetAllTodo 查询所有Todo
func GetAllTodo() (todoList []*Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	return
}
