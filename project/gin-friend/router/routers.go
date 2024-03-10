package router

import (
	"GOwork/project/gin-friend/controllers"

	"GOwork/project/gin-friend/config"
	"GOwork/project/gin-friend/pkg/logger"

	"github.com/gin-contrib/sessions"
	sessions_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// 返回值为路由引擎
func Router() *gin.Engine {
	//先把路由引擎的默认设置配置好
	r := gin.Default()
	//
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)
	//redis中间件
	store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	//想要在url为user中访问，所以定义分组，里面把访问方法都定义上
	user := r.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("/login", controllers.UserController{}.Login)

	}

	player := r.Group("/player")
	{
		player.POST("/list", controllers.PlayerController{}.GetPlayers)
	}

	vote := r.Group("/vote")
	{
		vote.POST("/add", controllers.VoteController{}.AddVote)
	}

	r.POST("/ranking", controllers.PlayerController{}.GetRanking)

	return r
}
