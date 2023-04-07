package blog

type Blog struct {
	ID int
	*CreateBlogRequest
}

type CreateBlogRequest struct {
	Title   string
	Content string
}
