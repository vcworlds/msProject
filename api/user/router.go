package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"msproject/router"
)

// RouterUser 是一个空的结构体
type RouterUser struct {
}

func init() {
	log.Print("init user router")
	router.Register(&RouterUser{})
}

// Route 方法实现了 Router 接口
// 该方法接受一个 *gin.Engine 参数 r，并将一个 POST 路由注册到 Gin 引擎
func (*RouterUser) Route(r *gin.Engine) {
	h := HandleFunc{}
	// 注册 POST 路由，当访问 /api/project/login/getCaptcha 时调用 h.GetCaptcha 方法
	r.GET("/api/project/login/getCaptcha", h.GetCaptcha)
}
