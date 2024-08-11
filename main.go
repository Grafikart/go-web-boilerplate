package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var IsDevMode = os.Getenv("APP_ENV") == "dev"
var VitePort = 3000

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Welcome to My Go Server</title>
			%s
		</head>
		<body>
			<h1>Hello, World!</h1>
			<p>This is a simple web server written in Go.</p>
		</body>
		</html>
	`, GetAssetsTags())
}

func PushHandler(w http.ResponseWriter, r *http.Request) {
	pushMessage("Hello world")
	w.WriteHeader(204)
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/sse", sseHandler)
	server.HandleFunc("/push", PushHandler)
	AddAssetsHandler(server)
	server.HandleFunc("/", homeHandler)
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
