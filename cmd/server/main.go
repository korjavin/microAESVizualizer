package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	port := flag.String("port", "8080", "port to serve on")
	directory := flag.String("dir", "static", "directory of static files")
	flag.Parse()

	absDir, err := filepath.Abs(*directory)
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.Dir(absDir))
	http.Handle("/", addWasmContentTypeHeader(fileServer))

	log.Printf("Serving %s on HTTP port: %s\n", absDir, *port)
	log.Printf("Open http://localhost:%s in your browser", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

// Add correct MIME type for WebAssembly files
func addWasmContentTypeHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if filepath.Ext(r.URL.Path) == ".wasm" {
			w.Header().Set("Content-Type", "application/wasm")
		}
		next.ServeHTTP(w, r)
	})
}
