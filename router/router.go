package router

import (
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/quincy0/live-ai/controller"
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
	}
}

func RegisterLiveMoney(g *gin.RouterGroup) {
	v1 := g.Group("v1")
	{
		v1.GET("/goods/list", controller.GoodsList)

		v1.GET("/product/tags", controller.ProductTagList)

		v1.POST("/chapter/create", controller.ChapterCreate)
		v1.GET("/chapter/list", controller.ChapterList)
		v1.GET("/chapter/info", controller.ChapterInfo)
		v1.POST("/paragraph/edit", controller.ParagraphEdit)
		v1.POST("/script/create", controller.ScriptCreate)
		v1.GET("/script/list", controller.ScriptList)
		v1.GET("/script/info", controller.ScriptInfo)
		v1.POST("/scene/edit", controller.SceneEdit)
	}
}
