package app

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"douyin/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	// 页码
	Page int `json:"page"`
	// 每页数量
	PageSize int `json:"page_size"`
	// 总行数
	TotalRows int `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToResponse(data map[string]interface{}) {
	responseData := gin.H{
		"status_code": 0,
		"status_msg":  "success",
	}
	if data != nil {
		for s, v := range data {
			responseData[s] = v
		}
	}

	r.Ctx.JSON(http.StatusOK, responseData)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error, data map[string]interface{}) {
	responseData := gin.H{
		"status_code": err.Code(), 
		"status_msg": err.Msg(),
	}
	for s, v := range data {
		responseData[s] = v
	}
	r.Ctx.JSON(http.StatusOK, responseData)

}