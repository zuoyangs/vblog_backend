package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zuoyangs/vblog_backend/blog/http"
	"github.com/zuoyangs/vblog_backend/blog/impl"
)

func main() {

	r := gin.Default()
	//把业务路由通过 web 框架，对外提供服务
	http.NewHandler(impl.NewImpl()).Registry(r)

}
