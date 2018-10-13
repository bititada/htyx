package router

import (
	"htyx/handler/article"
	"htyx/handler/audio"
	"htyx/handler/category"
	"htyx/handler/comments"
	"htyx/handler/likelist"
	"htyx/handler/play"
	"htyx/handler/sd"
	"htyx/handler/user"
	"htyx/handler/video"

	"htyx/router/middleware"
	"net/http"

	//"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	//"github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	//g.Use(gin.Recovery())
	g.Use(mw...)
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	//g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//pprof.Register(g)

	g.Static("/image", "./image")
	g.GET("/playstream/:streamname/:streamtype", play.PlayStream) //streamtype 音频audio视频video

	g.POST("/user/token", user.CreateToken)

	g.POST("/like", middleware.AuthMiddleware(), user.ChangeLikeStatus)
	g.POST("/likecomment", middleware.AuthMiddleware(), user.ChangeCommentLikeStatus)

	g.GET("/category/:cltype", category.GetCategory)
	g.GET("/article/list/:categoryid/:start", article.GetListByCateId)
	g.GET("/audio/list/:categoryid/:start", audio.GetListByCateId)
	g.GET("/video/list/:categoryid/:start", video.GetListByCateId)

	/*注意 /article/:categoryid/:start与/article/:id/:cltype会产生路由错误
	/article/list/:categoryid/:start 与 /article/:id/:cltype 加上一个list还是会产生路由错误
	只有/article/list/:categoryid/:start 与 /article/one/:id/:cltype list 和one才不会冲突
	*/

	ar := g.Group("/article")
	ar.Use(middleware.AuthMiddleware())
	{

		ar.GET("/one/:id/:cltype", article.GetOne)

	}
	au := g.Group("/audio")
	au.Use(middleware.AuthMiddleware())
	{
		au.GET("/one/:id/:cltype", audio.GetOne)

	}

	vi := g.Group("/video")
	vi.Use(middleware.AuthMiddleware())
	{
		vi.GET("/one/:id/:cltype", video.GetOne)

	}

	li := g.Group("/likelist")
	li.Use(middleware.AuthMiddleware())
	{

		li.GET("/:cltype/:start", likelist.GetListByCltype)
	}
	co := g.Group("/comments")
	co.Use(middleware.AuthMiddleware())
	{
		co.GET("/new/:clid/:cltype/:start", comments.GetNewComments)
		co.GET("/hot/:clid/:cltype/:start", comments.GetHotComments)
		co.POST("/insert", comments.CreateComment)
	}
	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}
	return g

}
