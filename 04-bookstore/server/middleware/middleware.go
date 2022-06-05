package middleware

import (
	"log"
	"mime"
	"net/http"
)

// 基本格式
func MiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// 处理自己的业务
		next.ServeHTTP(w, req)
	})
}

// Logging
// 输出每个到达的 HTTP 请求的一些概要信息
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("recv a %s request from %s\n", req.Method, req.RemoteAddr)
		next.ServeHTTP(w, req)
	})
}

// Validating
// 对每个 http 请求的头部进行检查，检查 Content-Type
// 头字段所表示的媒体类型是否为 application/json。
func Validating(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		contentType := req.Header.Get("Content-Type")
		mediatype, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if mediatype != "application/json" {
			http.Error(w, "invalid Content-Type", http.StatusUnsupportedMediaType)
			return
		}
		next.ServeHTTP(w, req)
	})
}
