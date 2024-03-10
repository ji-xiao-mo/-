package router

import (
	"GOwork/project/gin-ranking/controllers"
	"net/http"

	"GOwork/project/gin-ranking/pkg/logger"

	"github.com/gin-gonic/gin"
)

// 返回值为路由引擎
func Router() *gin.Engine {
	//先把路由引擎的默认设置配置好
	r := gin.Default()
	//
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)
	//想要在url为user中访问，所以定义分组，里面把访问方法都定义上
	user := r.Group("/user")
	{
		user.GET("/info/:id", controllers.UserController{}.GetUserInfo)

		user.POST("/list", controllers.UserController{}.GetList)
		user.POST("/add", controllers.UserController{}.AddUser)
		user.POST("/updata", controllers.UserController{}.UpdateUser)
		user.POST("/delete", controllers.UserController{}.DeleteUser)
		user.POST("/list/test", controllers.UserController{}.GetUserListTest)

		user.PUT("/add", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user add")
		})

		user.DELETE("/delete", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user delete")
		})
	}

	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}

	return r
}
