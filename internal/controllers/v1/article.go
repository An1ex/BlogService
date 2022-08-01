package v1

import (
	"strconv"

	"BlogService/global"
	"BlogService/internal/service"
	"BlogService/pkg/app"
	"BlogService/pkg/errcode"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

// @Summary 获取单篇文章
// @Produce  json
// @Param id path uint true "文章 ID"
// @Param title query string false "文章标题" maxlength(100)
// @Param state query uint false "状态" Enums(0, 1) default(1)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [get]
func (t Article) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	param := service.GetArticleRequest{
		ID: uint(id),
	}
	response := app.NewResponse(c)
	valid, verrs := app.BindAndValid(c, &param)
	if !valid { //入参校验或绑定参数失败
		global.Logger.WithFields(log.Fields{
			"error": verrs.Error(),
		}).Error("BindAndValid failed")
		errRsp := errcode.InvalidParams.WithDetails(verrs.Errors()...)
		response.ToErrorResponse(errRsp)
	} else { //参数校验和绑定参数成功
		svc := service.New(c.Request.Context())
		article, err := svc.GetArticle(&param)
		if err != nil { //获取文章失败
			global.Logger.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Get Article failed")
			errRsp := errcode.ErrorGetArticleFail.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			return
		}
		//获取标签列表成功
		response.ToResponse(article)
	}
}

// @Summary 获取多篇文章
// @Produce  json
// @Param title query string false "文章标题" maxlength(100)
// @Param state query uint false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.SwaggerArticle "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (t Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	valid, verrs := app.BindAndValid(c, &param)

	if !valid { //入参校验或绑定参数失败
		global.Logger.WithFields(log.Fields{
			"error": verrs.Error(),
		}).Error("BindAndValid failed")
		errRsp := errcode.InvalidParams.WithDetails(verrs.Errors()...)
		response.ToErrorResponse(errRsp)
	} else { //参数校验和绑定参数成功
		svc := service.New(c.Request.Context())

		totalRows, err := svc.CountArticle(&service.CountArticleRequest{
			Title: param.Title,
			State: param.State,
		})
		if err != nil { //统计文章数量失败
			global.Logger.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Count Article failed")
			response.ToErrorResponse(errcode.ErrorCountArticleFail)
			return
		}

		pager := app.Pager{
			Page:     app.GetPage(c),
			PageSize: app.GetPageSize(c),
		}
		tags, err := svc.GetArticleList(&param, &pager)
		if err != nil { //获取文章列表失败
			global.Logger.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Get Article List failed")
			errRsp := errcode.ErrorGetArticleListFail.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			return
		}
		//获取文章列表成功
		response.ToResponseList(tags, int(totalRows))
	}
}

// @Summary 新增文章
// @Produce  json
// @Param title body string true "文章标题" minlength(2) maxlength(100)
// @Param desc body string false "文章简介" maxlength(255)
// @Param content body string true "文章内容" minlength(2)
// @Param cover_image_url body string false "文章封面" minlength(2) maxlength(255)
// @Param created_by body string true "创建者" minlength(2) maxlength(100)
// @Param state body uint false "状态" Enums(0, 1) default(1)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (t Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, verrs := app.BindAndValid(c, &param)

	if !valid { //入参校验或绑定参数失败
		global.Logger.WithFields(log.Fields{
			"error": verrs.Error(),
		}).Error("BindAndValid failed")
		errRsp := errcode.InvalidParams.WithDetails(verrs.Errors()...)
		response.ToErrorResponse(errRsp)
	} else { //参数校验和绑定参数成功
		svc := service.New(c.Request.Context())
		err := svc.CreateArticle(&param)
		if err != nil { //创建文章失败
			global.Logger.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Create Article failed")
			errRsp := errcode.ErrorCreateArticleFail.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			return
		}
		//新增文章成功
		response.ToResponse(gin.H{})
	}
}

// @Summary 更新文章
// @Produce  json
// @Param id path uint true "文章 ID"
// @Param title body string false "文章标题" maxlength(100)
// @Param desc body string false "文章简介" maxlength(255)
// @Param content body string false "文章内容"
// @Param cover_image_url body string false "文章封面" maxlength(255)
// @Param state body uint true "状态" Enums(0, 1)
// @Param update_by body string true "修改者" minlength(2) maxlength(100)
// @Success 200 {array} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (t Article) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	param := service.UpdateArticleRequest{
		ID: uint(id),
	}
	response := app.NewResponse(c)
	valid, verrs := app.BindAndValid(c, &param)

	if !valid { //入参校验或绑定参数失败
		global.Logger.WithFields(log.Fields{
			"error": verrs.Error(),
		}).Error("BindAndValid failed")
		errRsp := errcode.InvalidParams.WithDetails(verrs.Errors()...)
		response.ToErrorResponse(errRsp)
	} else { //参数校验和绑定参数成功
		svc := service.New(c.Request.Context())
		err := svc.UpdateArticle(&param)
		if err != nil { //更新文章失败
			global.Logger.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Update Article failed")
			errRsp := errcode.ErrorUpdateArticleFail.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			return
		}
		//更新文章成功
		response.ToResponse(gin.H{})
	}
}

// @Summary 删除文章
// @Produce  json
// @Param id path uint true "文章 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (t Article) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	param := service.DeleteArticleRequest{
		ID: uint(id),
	}
	response := app.NewResponse(c)
	valid, verrs := app.BindAndValid(c, &param)

	if !valid { //入参校验或绑定参数失败
		global.Logger.WithFields(log.Fields{
			"error": verrs.Error(),
		}).Error("BindAndValid failed")
		errRsp := errcode.InvalidParams.WithDetails(verrs.Errors()...)
		response.ToErrorResponse(errRsp)
	} else { //参数校验和绑定参数成功
		svc := service.New(c.Request.Context())
		err := svc.DeleteArticle(&param)
		if err != nil {
			global.Logger.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Delete Article failed")
			errRsp := errcode.ErrorDeleteArticleFail.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			return
		}
		//删除标签成功
		response.ToResponse(gin.H{})
	}
}
