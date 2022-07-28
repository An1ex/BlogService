package v1

import (
	"strconv"

	"BlogService/global"
	"BlogService/internal/service"
	"BlogService/pkg/app"
	"BlogService/pkg/errcode"
	"BlogService/pkg/upload"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

// @Summary 上传文件
// @Produce  json
// @Param file formData file true "上传文件"
// @Param type body int true "文件格式" Enums(1, 2, 3)
// @Success 200 {object} service.FileInfo "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/upload/file [post]
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	fileHeader, err := c.FormFile("file")
	if err != nil { //获取上传文件失败
		errRsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	fileType, _ := strconv.Atoi(c.PostForm("type"))
	if fileHeader == nil || fileType <= 0 { //上传文件为空,入参不符
		errRsp := errcode.InvalidParams
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c)
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), fileHeader)
	if err != nil { //保存上传文件失败
		global.Logger.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Upload file failed")
		errRsp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errRsp)
		return
	}

	response.ToResponse(gin.H{
		"file_name":       fileInfo.Name,
		"file_access_url": fileInfo.AccessUrl,
	})
}
