package controllers

import (
	"GOwork/project/gin-ranking/models"
	"GOwork/project/gin-ranking/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	logger.Write("日志信息", "user")
	idStr := c.Param("id")
	name := c.Param("name")

	id, _ := strconv.Atoi(idStr)

	user, _ := models.GetUserTest(id)

	ReturnSuccess(c, 0, name, user, 2)
}

// 调用models中user.go的添加方法,在数据库中添加数据
func (u UserController) AddUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	id, err := models.AddUser(username)
	if err != nil {
		ReturnError(c, 4002, "保存错误")
	}
	ReturnSuccess(c, 0, "保存成功", id, 1)
}

// 通过指定id,来更新符合条件的username
func (u UserController) UpdateUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	models.UpdataUser(id, username)
	ReturnSuccess(c, 0, "更新成功", true, 1)
}

// 通过指定的id主键，来删除表格中的数据
func (u UserController) DeleteUser(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	err := models.DeleteUser(id)
	if err != nil {
		ReturnError(c, 400, "删除失败")
	}
	ReturnSuccess(c, 0, "删除成功", true, 1)
}

func (u UserController) GetList(c *gin.Context) {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("捕获异常", err)
	// 	}
	// }()  上面的错误捕获代码，被router中的中间件代替了
	num1 := 1
	num2 := 0
	num3 := num1 / num2
	ReturnError(c, 4004, num3)
}

func (u UserController) GetUserListTest(c *gin.Context) {
	users, err := models.GetUserListTest()
	if err != nil {
		ReturnError(c, 4004, "没有相关数据")
	}
	ReturnSuccess(c, 0, "获取成功", users, 1)
}
