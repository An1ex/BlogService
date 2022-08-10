package app

import (
	"net/http"

	"BlogService/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Response struct {
	ctx *gin.Context
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	TotalRows int `json:"totalRows"`
}

func NewResponse(c *gin.Context) *Response {
	return &Response{ctx: c}
}

func (r *Response) ToResponse(data interface{}) {
	r.ctx.JSON(http.StatusOK, gin.H{
		"code":    errcode.Success.Code,
		"msg":     errcode.Success.Msg,
		"details": data,
	})
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.ctx.JSON(http.StatusOK, gin.H{
		"code": errcode.Success.Code,
		"msg":  errcode.Success.Msg,
		"details": gin.H{
			"list": list,
			"pager": Pager{
				Page:      GetPage(r.ctx),
				PageSize:  GetPageSize(r.ctx),
				TotalRows: totalRows,
			},
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	r.ctx.JSON(err.StatusCode(), gin.H{
		"code":    err.Code,
		"msg":     err.Msg,
		"details": err.Details,
	})
}
