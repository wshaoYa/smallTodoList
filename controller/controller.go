package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"smallTodoList/models"
)

var (
	fail    = "fail"
	success = "success"
)

// InitIndexHtml 初始化index页面
func InitIndexHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// CreateTodo 增加一个Todo
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  fail,
			"data": err.Error(),
		})
		return
	}

	err = models.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  fail,
			"data": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  success,
			"data": todo,
		})
	}
}

// DeleteTodo 删除一个Todo
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	err := models.DeleteTodo(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  fail,
			"data": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  success,
			"data": fmt.Sprintf("id=%v 删除成功", id),
		})
	}
}

// UpdateTodoStatus 修改Todo状态
func UpdateTodoStatus(c *gin.Context) {
	id := c.Param("id")

	err := models.UpdateTodoStatus(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  fail,
			"data": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  success,
			"data": nil,
		})
	}
}

// GetAllTodo 查询所有Todo
func GetAllTodo(c *gin.Context) {
	todos, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  fail,
			"data": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, todos)
	}
}
