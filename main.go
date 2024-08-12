package main

import (
	"context"
	"embed"
	"fmt"
	"grafikart/boilerplate/server"
	"log"
	"net/http"
)

//go:embed public/assets/*
var assets embed.FS

func main() {
	viteAssets := server.NewViteAssets(assets)
	frontMiddleware := createFrontEndMiddleware(*viteAssets)
	http.HandleFunc("/sse", server.SSEHandler)
	http.HandleFunc("/assets/", viteAssets.ServeAssets)
	http.HandleFunc("/", frontMiddleware(server.HomeHandler))
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createFrontEndMiddleware(vite server.ViteAssets) func(func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	html := vite.GetHeadHTML()
	return func(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "assets", html)
			next(w, r.WithContext(ctx))
		}
	}
}
