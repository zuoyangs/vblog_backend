package impl

import (
	"context"

	"github.com/zuoyangs/vblog_backend/blog"
)

func NewImpl() *Impl {
	return &Impl{}
}

type Impl struct{}

func (i *Impl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {

	return nil, nil
}
