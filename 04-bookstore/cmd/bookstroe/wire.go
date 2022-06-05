//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	storegrpc "go_advanced_training/04-bookstore/internal/store"
	server2 "go_advanced_training/04-bookstore/server"
)

//grpcstore := storegrpc.NewStore(":8091")
////NewBookStoreServer(addr string, s store.Store) *BookStoreServer
//srv := server2.NewBookStoreServer("8092", grpcstore)

func NewServer() (*BookServer, error) {
	panic(wire.Build(storegrpc.Provider, server2.Provider, NewBookServer))
}
