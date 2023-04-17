package impl

import (
	"context"

	"github.com/zuoyangs/vblog_backend/blog"
)

// 构造函数，返回一个 Impl 实例
func NewImpl() *Impl {
	return &Impl{}
}

// Impl 结构体
type Impl struct{}

// CreateBlog 方法，接收一个 ctx 和 req 参数，返回 blog 和 error
func (i *Impl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {

	return nil, nil
}
