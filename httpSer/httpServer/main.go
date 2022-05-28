package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var hostAndPort = ":8090"

func main() {
	// 注册路由
	http.HandleFunc("/", hello)
	http.HandleFunc("/echo", echo)

	//
	ctx, cancel := context.WithCancel(context.Background())
	group, errCtx := errgroup.WithContext(ctx)

	fmt.Println("Listen on " + hostAndPort)
	server := &http.Server{Addr: hostAndPort}

	// 启动服务
	group.Go(func() error {
		return startHttpServer(server)
	})
	//
	group.Go(func() error {
		<-errCtx.Done() //阻塞。因为 cancel、timeout、deadline 都可能导致 Done 被 close
		fmt.Println("Get err context done channel, http server will stop")
		// 关闭 http server
		return server.Shutdown(errCtx)
	})

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	group.Go(func() error {
		for {
			select {
			case sig := <-sigs:
				fmt.Printf("Get signal %v\n", sig)
				cancel()
			case <-errCtx.Done():
				fmt.Printf("Err context done! \n")
				return errCtx.Err()
			}
		}
	})
	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("all group done!")

}

func startHttpServer(ser *http.Server) error {
	err := ser.ListenAndServe()
	return err
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello go")
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello go echo")
}
