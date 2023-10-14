package main

import (
	"smallTodoList/dao"
	"smallTodoList/models"
	"smallTodoList/routers"
)

func main() {
	// 初始化数据库
	dao.InitMySQL()

	// 模型绑定
	err := dao.DB.AutoMigrate(&models.Todo{})
	if err != nil {
		panic(err)
	}

	// 注册路由
	routers.SetupRouter()
}
