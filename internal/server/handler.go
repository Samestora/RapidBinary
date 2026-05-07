package server

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"rapidbinary/ui"
	"time"
)

func RunServer(port int, serveDir, uploadDir string, canUpload, useLog bool) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && canUpload {
			handleUpload(w, r, uploadDir)
			return
		}

		if canUpload && r.URL.Path == "/upload" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(ui.UploadHTML)
			return
		}

		fullPath := filepath.Join(serveDir, filepath.Clean(r.URL.Path))
		info, err := os.Stat(fullPath)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		if info.IsDir() {
			RenderFolder(w, fullPath, r.URL.Path, canUpload)
			return
		}
		http.ServeFile(w, r, fullPath)
	})

	var handler http.Handler = mux
	if useLog {
		handler = logger(mux)
	}

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("RapidBinary serving [%s] on http://localhost%s\n", serveDir, addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s from %s (%s)", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}

func handleUpload(w http.ResponseWriter, r *http.Request, dir string) {
	r.ParseMultipartForm(32 << 20)
	file, h, err := r.FormFile("file")
	if err != nil {
		return
	}
	defer file.Close()

	os.MkdirAll(dir, 0755)
	dst, _ := os.Create(filepath.Join(dir, h.Filename))
	defer dst.Close()
	io.Copy(dst, file)

	fmt.Fprintf(w, "<h3>Received %s</h3><a href='/upload'>Back</a>", h.Filename)
}
