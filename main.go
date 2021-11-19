package main

import (
	"github.com/qiuyuhome/go-gin-blog-api/internal/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewRouter()

	s := &http.Server{
		Addr:           ":80",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
