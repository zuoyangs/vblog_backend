package tag

import "context"

type Service interface {
	Query(ctx context.Context, req *QueryRequest) (*TagSet, error)
}


type QueryRequest struct{
	BlogId int
}