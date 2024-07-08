package user

import (
	"github.com/gin-gonic/gin"
	"msproject/common"
	"net/http"
)

type HandleFunc struct {
}

func (HandleFunc) GetCaptcha(ctx *gin.Context) {
	rsp := common.Response{}
	ctx.JSON(http.StatusOK, rsp.Success("获取成功", "123456"))
}
