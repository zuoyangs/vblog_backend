package blog

import "context"

type Service interface {
	CreateBlog(context.Context, *CreateBlogRequest) (Blog, error)
}
