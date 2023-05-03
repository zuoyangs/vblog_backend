package blog

import "context"

// Service 接口的定义
type Service interface {

	// 创建文章
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)

	// 更新文章
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)

	//文章删除
	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)

	//文章列表
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)

	//文章详情
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)

	//更新文章的状态
	UpdateBlogStatus(context.Context, *UpdateBlogStatusRequest) (*Blog, error)
}
