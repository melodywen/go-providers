package implement_servers

import (
	"context"
	"fmt"
	"github.com/melodywen/go-providers/servers/gen/golang"
	"math/rand"
)

type BlogImplementServer struct {
}

func NewBlogImplementServer() *BlogImplementServer {
	return &BlogImplementServer{}
}

func (b BlogImplementServer) Index(ctx context.Context, request *golang.BlogIndexRequest) (*golang.BlogIndexResponse, error) {
	fmt.Println(1213131313)
	return &golang.BlogIndexResponse{Status: int32(rand.Int())}, nil
}
