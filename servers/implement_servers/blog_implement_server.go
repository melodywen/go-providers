package implement_servers

import (
	"context"
	"github.com/melodywen/go-providers/servers/gen/golang"
	"github.com/melodywen/go-providers/servers/gen/golang/model"
	"math/rand"
)

type BlogImplementServer struct {
}

func (b BlogImplementServer) Store(ctx context.Context, request *golang.BlogCreateRequest) (*golang.BlogIndexResponse, error) {
	panic("implement me")
}

func (b BlogImplementServer) Update(ctx context.Context, request *golang.BlogIndexRequest) (*golang.BlogIndexResponse, error) {
	panic("implement me")
}

func (b BlogImplementServer) Show(ctx context.Context, request *golang.BlogIndexRequest) (*golang.BlogIndexResponse, error) {
	panic("implement me")
}

func (b BlogImplementServer) Delete(ctx context.Context, request *golang.BlogIndexRequest) (*golang.BlogIndexResponse, error) {
	panic("implement me")
}

func NewBlogImplementServer() *BlogImplementServer {
	return &BlogImplementServer{}
}

func mockBlogModel(size int) (response []*model.BlogModel) {
	for i := 0; i < size; i++ {
		response = append(response, &model.BlogModel{
			Id:       int32(rand.Int()),
			Title:    "",
			Describe: "",
			Author:   "",
			Status:   2,
		})
	}
	return response
}

func (b BlogImplementServer) Index(ctx context.Context, request *golang.BlogIndexRequest) (*golang.BlogIndexResponse, error) {
	data := mockBlogModel(int(request.Size))
	return &golang.BlogIndexResponse{Data: data}, nil
}
