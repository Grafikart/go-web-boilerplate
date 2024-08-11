package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed public/assets/*
var assets embed.FS

// Add /assets handler on http
func AddAssetsHandler(server *http.ServeMux) {
	// Proxy everything to vite in dev mode
	if IsDevMode {
		server.HandleFunc("/assets/", redirectToVite)
		return
	}

	// Otherwise serve static assets from public directory
	assetsFs, err := fs.Sub(assets, "public")
	if err != nil {
		panic(fmt.Sprintf("Cannot sub public directory %v", err))
	}
	server.Handle("/assets/", http.FileServer(http.FS(assetsFs)))
}

func GetAssetsTags() string {
	if IsDevMode {
		return `<script type="module" src="http://localhost:3000/@vite/client"></script>
			<script src="http://localhost:3000/assets/main.tsx" type="module"></script>`
	}

	return `<script src="/assets/index.js" type="module"></script><link rel="stylesheet" href="/assets/index.css">`
}

func redirectToVite(w http.ResponseWriter, r *http.Request) {
	u := *r.URL
	u.Host = strings.Split(r.Host, ":")[0] + ":3000"
	w.Header().Set("Location", u.String())
	w.WriteHeader(301)
}
