package routers

import (
	"github.com/gin-gonic/gin"
	"smallTodoList/controller"
)

// SetupRouter 注册路由
func SetupRouter() {
	r := gin.Default()

	// 加载初始网页index
	r.Static("/static", "statics")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.InitIndexHtml)

	v1Group := r.Group("/v1")
	{
		//增
		v1Group.POST("/todo", controller.CreateTodo)
		//删
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
		//改
		v1Group.PUT("/todo/:id", controller.UpdateTodoStatus)
		//查
		v1Group.GET("/todo", controller.GetAllTodo)
	}

	err := r.Run(":9090")
	if err != nil {
		panic(err)
	}
}
