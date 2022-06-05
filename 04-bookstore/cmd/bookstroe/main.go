package main

import (
	"context"
	storegrpc "go_advanced_training/04-bookstore/internal/store"
	"go_advanced_training/04-bookstore/server"
	"go_advanced_training/04-bookstore/store"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type BookServer struct { // 最终需要的对象
	S       storegrpc.GrpcStore
	HttsSer *server.BookStoreServer
}

func NewBookServer() (store.Store, *server.BookStoreServer) {
	s := storegrpc.GrpcStore{}
	httpSer := &server.BookStoreServer{}
	return s, httpSer
}

func main() {
	//grpcstore := storegrpc.NewStore(":8091")
	////NewBookStoreServer(addr string, s store.Store) *BookStoreServer
	//srv := server2.NewBookStoreServer("8092", grpcstore)

	grpcstore, srv := NewServer()
	defer grpcstore.Conn.Close()
	defer grpcstore.Cancel()
	errChan, err := srv.ListenAndServer()
	if err != nil {
		log.Println("web server start failed:", err)
		return
	}
	log.Println("web server start ok listen on 8092")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case err = <-errChan:
		log.Println("web server run failed: ", err)
		return
	case <-c:
		log.Println("bookstore program is exiting...")
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = srv.Shutdown(ctx)
	}

	if err != nil {
		log.Println("bookstore program exit error:", err)
		return
	}
	log.Println("bookstore program exit ok")
}
