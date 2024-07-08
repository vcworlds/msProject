package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandleFunc struct {
}

func (HandleFunc) GetCaptcha(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "success")
}
