package implement_servers

import (
	"context"
	"fmt"
	"github.com/melodywen/go-providers/servers/gen/golang"
	"github.com/melodywen/go-providers/servers/gen/golang/model"
	"math/rand"
)

type BlogImplementServer struct {
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
	return &golang.BlogIndexResponse{Items: data}, nil
}

func (b BlogImplementServer) Store(ctx context.Context, form *model.BlogForm) (*model.BlogModel, error) {
	fmt.Println(form)
	data := mockBlogModel(1)
	blog := data[0]
	blog.Title = form.Title
	blog.Describe = form.Describe
	blog.Author = form.Author
	return data[0], nil
}

func (b BlogImplementServer) Update(ctx context.Context, form *model.BlogForm) (*model.BlogModel, error) {
	fmt.Println(form)
	data := mockBlogModel(1)
	blog := data[0]
	blog.Id = form.Id
	blog.Title = form.Title
	blog.Describe = form.Describe
	blog.Author = form.Author
	return data[0], nil
}