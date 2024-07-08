package main

import (
	"github.com/gin-gonic/gin"
	"msproject/common"
	"msproject/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	common.Run(r, "webcenter", ":8045")

}
