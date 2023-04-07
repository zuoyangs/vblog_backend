package blog

import "context"

// Service 接口的定义
type Service interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
}
