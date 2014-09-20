package main

import "github.com/gohttp/logger"
import "github.com/gohttp/app"
import "net/http"
import "io"

func main() {
	a := app.New()

	a.Use(logger.New())

	a.Get("/", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("hello"))
		res.Write([]byte(" world"))
	}))

	a.Get("/yahoo", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		page, _ := http.Get("http://yahoo.com")
		defer page.Body.Close()
		io.Copy(res, page.Body)
	}))

	a.Listen(":3000")
}
