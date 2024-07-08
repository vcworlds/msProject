package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"msproject/common"
	"msproject/pkg/e"
	"net/http"
	"time"
)

type HandleFunc struct {
}

func (HandleFunc) GetCaptcha(ctx *gin.Context) {
	rsp := common.Response{}
	// 获取数据
	mobile := ctx.PostForm("mobile")
	// 数据校验
	if !common.VerifyMobile(mobile) {
		ctx.JSON(http.StatusOK, rsp.Fail("手机号接收错误", e.MobileValidate))
		return
	}
	// 随机生成验证码
	code := "1234"
	// go协程调用第三方平台获取验证码
	go func() {
		time.Sleep(2 * time.Second)
		log.Print("短信平台调用成功，正在发送")
		log.Printf("将手机号和验证码存入redis成功register:%s:%s", mobile, code)
	}()
	ctx.JSON(http.StatusOK, rsp.Success("获取成功", "123456"))
}
