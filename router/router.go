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
