package store

import (
	"context"
	"github.com/google/wire"
	v1 "go_advanced_training/04-bookstore/api/bookstore/v1"
	"go_advanced_training/04-bookstore/store"
	"google.golang.org/grpc"
	"log"
	"time"
)

type GrpcStore struct {
	Conn     *grpc.ClientConn
	storeCli v1.BookstoreClient
	Ctx      context.Context
	Cancel   context.CancelFunc
}

var Provider = wire.NewSet(NewStore)

func NewStore() *GrpcStore {
	//addr string
	conn, err := grpc.Dial("8091", grpc.WithInsecure())
	//defer conn.Close()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	storeCli := v1.NewBookstoreClient(conn)
	// 需要close conn
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// defer cancel()

	return &GrpcStore{
		Conn:     conn,
		storeCli: storeCli,
		Ctx:      ctx,
		Cancel:   cancel,
	}
}

// 调用具体的GRPC方法....
func (s *GrpcStore) Create(book *store.Book) error {

	//userIndexReponse, err := userClient.UserIndex(ctx, &user.UserIndexRequest{Page: 1, PageSize: 12})
	//if err != nil {
	//	log.Printf("user index could not greet: %v", err)
	//}
	booReq := v1.BookRequest{}
	ctx := context.Background()
	_, err := s.storeCli.CreateBook(ctx, &booReq)
	if err != nil {
		return err
	}

	return nil
}

func (s *GrpcStore) Update(*store.Book) error {
	//TODO
	return nil
}

func (s *GrpcStore) GetAll() ([]store.Book, error) {
	//TODO
	return nil, nil
}
