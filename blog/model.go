package blog

type Blog struct {
	//文章ID
	Id int
	//文章摘要信息，通过提取 Content 内容和获取
	Sumary string
	//创建时间
	CreateAt int64
	//发布时间
	PublishAt int64
	//用户提交数据
	*CreateBlogRequest
	//文章状态
	status Status
}

type BlogSet struct {
	//列表数据
	Items []*Blog
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{}
}

type CreateBlogRequest struct {
	//文章图片
	TitleImg string
	//文章标题
	TitleName string
	//文章子标题
	SubTitle string
	//文章内容
	Content string
	//文章作者
	Author string
}

func NewPutUpdateBlogRequest() *UpdateBlogRequest {
	return &UpdateBlogRequest{
		UpdateMode: UPDATE_MODE_PUT,
	}
}
func NewPatchUpdateBlogRequest() *UpdateBlogRequest {
	return &UpdateBlogRequest{
		UpdateMode: UPDATE_MODE_PATCH,
	}
}

type UpdateBlogRequest struct {
	UpdateMode UpdateMode
	*CreateBlogRequest
}

type DeleteBlogRequest struct {
	Id int
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

type QueryBlogRequest struct {
	PageSize   int
	PageNumber int
	Keywords   string
}

type DescribeBlogRequest struct {
	Id int
}

type UpdateBlogStatusRequest struct {
	//文章ID
	Id     int
	status Status
}
