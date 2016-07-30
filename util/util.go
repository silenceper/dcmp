package util

import "github.com/gin-gonic/gin"

type result struct {
	Data    interface{} `json:"data"`
	RetCode int32       `json:"retCode"`
	RetMsg  string      `json:"retMsg"`
}

func RenderSuccess(ctx *gin.Context, obj interface{}) {
	res := &result{}
	res.Data = obj
	res.RetCode = 0
	res.RetMsg = ""
	ctx.JSON(200, res)
}

func RenderError(ctx *gin.Context, retCode int32, retMsg string) {
	res := &result{}
	res.RetCode = retCode
	res.RetMsg = retMsg
	ctx.JSON(200, res)
}
