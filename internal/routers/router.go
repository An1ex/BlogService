package routers

import (
	"net/http"
	"time"

	_ "BlogService/docs"
	"BlogService/global"
	v1 "BlogService/internal/controllers/v1"
	"BlogService/internal/middleware"
	"BlogService/pkg/limiter"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var buckets []limiter.BucketRule

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.Config.Server.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	initBuckets()
	methodLimiter := limiter.NewMethodLimiter().AddBuckets(buckets...)
	r.Use(middleware.Limiter(methodLimiter))

	r.Use(middleware.Tracing())
	r.Use(middleware.ContextTimeout())
	r.Use(middleware.Translation())

	t := v1.NewTag()
	a := v1.NewArticle()
	f := v1.NewUpload()

	//注册Swagger接口路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//注册静态文件路由
	r.StaticFS("/static", http.Dir(global.Config.App.UploadSavePath))
	//注册鉴权路由
	r.GET("/auth", v1.GetAuth)

	//注册博客功能路由
	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())
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

		apiV1.POST("/articles_tag/:id", a.AppendTag)
		apiV1.DELETE("/articles_tag/:id", a.RemoveTag)

		apiV1.POST("/upload/file", f.UploadFile)
	}
	return r
}

//读取配置文件，初始化令牌桶
func initBuckets() {
	for _, lv := range global.Config.Limiter {
		buckets = append(buckets, limiter.BucketRule{
			Key:          lv.Key,
			FillInterval: lv.FillInterval * time.Second,
			Capacity:     lv.Capacity,
			Quantum:      lv.Quantum,
		})
	}
}
