package router

import (
	"github.com/gin-gonic/gin"
	"msproject/api/user"
)

// Router 接口定义，需要实现 Route 方法
type Router interface {
	Route(r *gin.Engine)
}

// RegisterRouter 是一个空的结构体
type RegisterRouter struct {
}

// New 函数创建并返回一个 RegisterRouter 实例
func New() RegisterRouter {
	return RegisterRouter{}
}

// Route 方法接受一个实现了 Router 接口的类型 ro 和一个 *gin.Engine 参数 r
// 并调用 ro 的 Route 方法，将 r 传递给它
func (RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

// InitRouter 函数接受一个 *gin.Engine 参数 r
// 创建一个 RegisterRouter 实例，并调用其 Route 方法，将 user.RouterUser{} 实例和 r 传递给它
func InitRouter(r *gin.Engine) {
	rg := New()
	rg.Route(&user.RouterUser{}, r)
}
