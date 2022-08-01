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

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取单个标签
// @Produce  json
// @Param id path uint true "标签 ID"
// @Param name query string false "标签名称" maxlength(100)
// @Param state query uint false "状态" Enums(0, 1) default(1)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [get]
func (t Tag) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	param := service.GetTagRequest{
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
		tag, err := svc.GetTag(&param)
		if err != nil { //获取标签失败
			global.Logger.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Get Tag failed")
			errRsp := errcode.ErrorGetTagFail.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			return
		}
		//获取标签列表成功
		response.ToResponse(tag)
	}
}

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query uint false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.SwaggerTag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
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

		totalRows, err := svc.CountTag(&service.CountTagRequest{
			Name:  param.Name,
			State: param.State,
		})
		if err != nil { //统计标签数量失败
			global.Logger.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Count Tag failed")
			response.ToErrorResponse(errcode.ErrorCountTagFail)
			return
		}

		pager := app.Pager{
			Page:     app.GetPage(c),
			PageSize: app.GetPageSize(c),
		}
		tags, err := svc.GetTagList(&param, &pager)
		if err != nil { //获取标签列表失败
			global.Logger.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Get Tag List failed")
			errRsp := errcode.ErrorGetTagListFail.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			return
		}
		//获取标签列表成功
		response.ToResponseList(tags, int(totalRows))
	}
}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(2) maxlength(100)
// @Param state body uint false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(2) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
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
		err := svc.CreateTag(&param)
		if err != nil { //创建标签失败
			global.Logger.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Create Tag failed")
			errRsp := errcode.ErrorCreateTagFail.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			return
		}
		//新增标签成功
		response.ToResponse(gin.H{})
	}
}

// @Summary 更新标签
// @Produce  json
// @Param id path uint true "标签 ID"
// @Param name body string false "标签名称" maxlength(100)
// @Param state body uint true "状态" Enums(0, 1)
// @Param update_by body string true "修改者" minlength(2) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	param := service.UpdateTagRequest{
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
		err := svc.UpdateTag(&param)
		if err != nil { //更新标签失败
			global.Logger.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Update Tag failed")
			errRsp := errcode.ErrorUpdateTagFail.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			return
		}
		//更新标签成功
		response.ToResponse(gin.H{})
	}
}

// @Summary 删除标签
// @Produce  json
// @Param id path uint true "标签 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	param := service.DeleteTagRequest{
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
		err := svc.DeleteTag(&param)
		if err != nil {
			global.Logger.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Delete Tag failed")
			errRsp := errcode.ErrorDeleteTagFail.WithDetails(err.Error())
			response.ToErrorResponse(errRsp)
			return
		}
		//删除标签成功
		response.ToResponse(gin.H{})
	}
}
