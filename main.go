package main

import (
	"context"
	"embed"
	"fmt"
	"grafikart/boilerplate/server"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:public
var assets embed.FS

func main() {
	publicFS, err := fs.Sub(assets, "public")
	if err != nil {
		panic(fmt.Sprintf("Cannot sub public directory from %v", err))
	}

	viteAssets := server.NewViteAssets(publicFS)
	frontMiddleware := createFrontEndMiddleware(*viteAssets)
	publicServer := http.FileServer(http.FS(publicFS))

	// Static Assets
	http.HandleFunc("/sse", server.SSEHandler)
	http.HandleFunc("/assets/", viteAssets.ServeAssets)

	// FrontEnd URLs
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve the root
		if r.URL.Path == "/" {
			frontMiddleware(server.HomeHandler)(w, r)
			return
		}
		// Otherwise serve public files
		publicServer.ServeHTTP(w, r)
	})
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
