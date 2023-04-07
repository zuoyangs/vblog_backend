package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zuoyangs/vblog_backend/blog"
)

func NewHandler(svr blog.Service) *Handler {
	return &Handler{
		service: svr,
	}
}

type Handler struct {
	service blog.Service
}

// 把该业务的 handler 注册给跟路由
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/vblog/api/v1/blogs", h.CreateBlog)
}

// http 协议逻辑处理
func (h *Handler) CreateBlog(ctx *gin.Context) {
	h.service.CreateBlog(nil, nil)
}
