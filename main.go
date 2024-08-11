package main

import (
	"embed"
	"fmt"
	"grafikart/boilerplate/server"
	"grafikart/boilerplate/utils"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed public/assets/*
var assets embed.FS
var IsDevMode = os.Getenv("APP_ENV") == "dev"
var VitePort = 3000

func main() {
	s := http.NewServeMux()
	s.HandleFunc("/sse", server.SSEHandler)
	server.AddAssetsHandler(s, utils.Force(fs.Sub(assets, "public")))
	s.HandleFunc("/", server.HomeHandler)
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", s))
}
