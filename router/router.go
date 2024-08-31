package router

import (
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/quincy0/live-ai/controller"
	"github.com/quincy0/live-ai/util"
	"github.com/quincy0/qpro/middleware"
)

func Init() *gin.Engine {
	r := gin.New()
	middleware.InitMiddleware(r)
	audio := r.Group("dh")
	RegisterAudio(audio)
	liveApi := r.Group("api/welive")
	RegisterLiveMoney(liveApi)
	return r

}

func RegisterAudio(g *gin.RouterGroup) {
	v1 := g.Group("v1")
	{
		v1.POST(
			"/audio/create",
			timeout.New(
				timeout.WithTimeout(10*time.Minute),
				timeout.WithHandler(func(c *gin.Context) {
					c.Next()
				}),
			),
			controller.AudioCreate,
		)
		v1.GET("/audio/list", controller.AudioList)
		v1.POST("/audio/notify/:sum", controller.AudioNotify)
	}
}

func RegisterLiveMoney(g *gin.RouterGroup) {
	v1 := g.Group("v1")
	{
		v1.POST("/register", controller.Register)
		v1.POST("/login", util.JWTAuth.LoginHandler)
		v1.GET("/refresh_token", util.JWTAuth.MiddlewareFunc(), util.JWTAuth.RefreshHandler)
		v1.GET("/hello", util.JWTAuth.MiddlewareFunc(), controller.Hello)

		v1.GET("/goods/list", controller.GoodsList)
		v1.GET("/product/tags", controller.ProductTagList)
		v1.GET("/script/tags", controller.ScriptTagList)
		v1.GET("/template/list", controller.RoomTemplateList)

		v1.POST("/chapter/create", util.JWTAuth.MiddlewareFunc(), controller.ChapterCreate)
		v1.GET("/chapter/list", util.JWTAuth.MiddlewareFunc(), controller.ChapterList)
		v1.GET("/chapter/info", util.JWTAuth.MiddlewareFunc(), controller.ChapterInfo)
		v1.POST("/paragraph/edit", util.JWTAuth.MiddlewareFunc(), controller.ParagraphEdit)

		v1.POST("/script/create", util.JWTAuth.MiddlewareFunc(), controller.ScriptCreate)
		v1.GET("/script/list", util.JWTAuth.MiddlewareFunc(), controller.ScriptList)
		v1.GET("/script/info", util.JWTAuth.MiddlewareFunc(), controller.ScriptInfo)
		v1.POST("/scene/edit", util.JWTAuth.MiddlewareFunc(), controller.SceneEdit)

		v1.POST("/room/create", util.JWTAuth.MiddlewareFunc(), controller.RoomCreate)
		v1.GET("/room/list", util.JWTAuth.MiddlewareFunc(), controller.RoomList)
		v1.GET("/room/info", util.JWTAuth.MiddlewareFunc(), controller.RoomInfo)
		v1.GET("/room/delete", util.JWTAuth.MiddlewareFunc(), controller.RoomDelete)

		v1.POST("/timbre/create", util.JWTAuth.MiddlewareFunc(), controller.TimbreCreate)
		v1.GET("/timbre/list", util.JWTAuth.MiddlewareFunc(), controller.TimbreList)

		v1.POST("/card/add", util.JWTAuth.MiddlewareFunc(), controller.AddCard)
		v1.POST("/card/delete", util.JWTAuth.MiddlewareFunc(), controller.DeleteCard)

		v1.POST("/window/get", util.JWTAuth.MiddlewareFunc(), controller.GetWindow)
		v1.POST("/window/add", util.JWTAuth.MiddlewareFunc(), controller.AddWindow)
		v1.POST("/window/delete", util.JWTAuth.MiddlewareFunc(), controller.DeleteWindow)
	}
}
