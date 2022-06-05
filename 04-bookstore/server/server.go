package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"go_advanced_training/04-bookstore/server/middleware"
	"go_advanced_training/04-bookstore/store"
	"io"
	"net/http"
	"time"
)

// BookStoreServer
// 接受请求, 并作出对应的数据操作
// 将http Server和 store做组合
type BookStoreServer struct {
	S store.Store

	Srv *http.Server
}

var Provider = wire.NewSet(NewBookStoreServer)

func NewBookStoreServer(s store.Store) *BookStoreServer {
	//addr string,
	srv := &BookStoreServer{
		s: s,
		srv: &http.Server{
			Addr: "8092",
		},
	}
	// 注册路由
	router := mux.NewRouter()
	router.HandleFunc("/book", srv.createBookHandler).Methods("POST")
	router.HandleFunc("/book/{id}", srv.updateBookHandler).Methods("POST")
	//router.HandleFunc("/book/{id}", srv.getBookHandler).Methods("GET")
	router.HandleFunc("/book", srv.getAllBookHandler).Methods("GET")
	//router.HandleFunc("/book/{id}", srv.getAllBookHandler).Methods("DELETE")
	//

	// 一定要绑定这个
	srv.srv.Handler = middleware.Logging(middleware.Validating(router))
	return srv
}

func (bs *BookStoreServer) ListenAndServer() (<-chan error, error) {
	var err error
	errChan := make(chan error, 1)
	go func() {
		err = bs.srv.ListenAndServe()
		errChan <- err
	}()

	select {
	case err = <-errChan:
		return nil, err
	case <-time.After(time.Second):
		return errChan, nil
	}
}

func (bs *BookStoreServer) Shutdown(ctx context.Context) error {
	return bs.srv.Shutdown(ctx)
}

func (bs *BookStoreServer) createBookHandler(w http.ResponseWriter,
	r *http.Request) {
	// 需要从POST中读取出Book的信息. 信息是json化的
	//dec := json.NewDecoder(r.Body)
	//var book store.Book
	//if err := dec.Decode(&book); err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	fmt.Println("api createBookHandler")
	book, err := getBookFromBody(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := bs.s.Create(book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("add succeed"))
}

func (bs *BookStoreServer) updateBookHandler(w http.ResponseWriter,
	r *http.Request) {
	//dec := json.NewDecoder(r.Body)
	//var book store.Book
	//if err := dec.Decode(&book); err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	fmt.Println("api updateBookHandler")
	book, err := getBookFromBody(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = bs.s.Update(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("up date succeed"))
}

//func (bs *BookStoreServer) getBookHandler(w http.ResponseWriter,
//	r *http.Request) {
//	fmt.Println("api getBookHandler")
//	// 取出id 然后调用s.get
//	id, ok := mux.Vars(r)["id"]
//	if !ok {
//		http.Error(w, "id not found in request", http.StatusBadRequest)
//		return
//	}
//	book, err := bs.s.Get(id)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	// book做json序列化, 发送回去
//	response(w, book)
//}

func (bs *BookStoreServer) getAllBookHandler(w http.ResponseWriter,
	r *http.Request) {
	fmt.Println("api getAllBookHandler")
	books, err := bs.s.GetAll()
	if err != nil {
		http.Error(w, "id not found in request", http.StatusBadRequest)
		return
	}
	response(w, books)
}

//func (bs *BookStoreServer) delBookHandler(w http.ResponseWriter,
//	r *http.Request) {
//	fmt.Println("api delBookHandler")
//	id, ok := mux.Vars(r)["id"]
//	if !ok {
//		http.Error(w, "id not found in request", http.StatusBadRequest)
//		return
//	}
//	err := bs.s.Delete(id)
//	if err != nil {
//		http.Error(w, "id not found in request", http.StatusInternalServerError)
//		return
//	}
//	w.Write([]byte("delete succeed"))
//}

func getBookFromBody(body io.ReadCloser) (*store.Book, error) {
	dec := json.NewDecoder(body)
	var book store.Book
	if err := dec.Decode(&book); err != nil {
		return nil, err
	}
	return &book, nil
}

func response(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
