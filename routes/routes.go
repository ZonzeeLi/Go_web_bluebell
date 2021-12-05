package routes

import (
	"bluebell/controller"
	_ "bluebell/docs"
	"bluebell/logger"
	"bluebell/middleware"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.RateLimitMiddleware(time.Second*2, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")

	// 注册
	v1.POST("signup", controller.SignupHandler)
	// 登录
	v1.POST("login", controller.LoginHandler)

	v1.GET("/post/:id", controller.GetPostDetailHandler)
	v1.GET("/posts", controller.GetPostListHandler)
	// 根据时间或分数获取帖子列表
	v1.GET("/posts2", controller.GetPostListHandler2)
	v1.GET("/community", controller.CommunityHandler)
	v1.GET("/community/:id", controller.CommunityDetailHandler)

	v1.Use(middleware.JWTAuthMiddleware()) // 应用JWT认证中间件

	{
		// 发帖
		v1.POST("/post", controller.CreatePostHandler)
		// 投票
		v1.POST("/vote", controller.PostVoteController)
	}

	pprof.Register(r) // 注册pprof相关路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
