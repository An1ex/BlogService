package routers

import (
	v1 "BlogService/internal/controllers/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	t := v1.NewTag()
	a := v1.NewArticle()
	apiV1 := r.Group("api/v1")
	{
		apiV1.POST("/tags", t.Create)
		apiV1.DELETE("/tags/:id", t.Delete)
		apiV1.PUT("/tags/:id", t.Update)
		apiV1.PATCH("/tags/:id/state", t.Update)
		apiV1.GET("/tags/:id", t.Get)
		apiV1.GET("/tags", t.List)

		apiV1.POST("/articles", a.Create)
		apiV1.DELETE("/articles/:id", a.Delete)
		apiV1.PUT("/articles/:id", a.Update)
		apiV1.PATCH("articles/:id/state", a.Update)
		apiV1.GET("/articles/:id", a.Get)
		apiV1.GET("/articles", a.List)
	}
	return r
}
