package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func RenderFolder(w http.ResponseWriter, localPath, urlPath string, canUpload bool) {
	files, _ := os.ReadDir(localPath)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprint(w, "<html><head><style>body{font-family:sans-serif;background:#f8f9fa;padding:40px;}.box{max-width:800px;margin:auto;background:#fff;padding:20px;border-radius:12px;box-shadow:0 4px 12px rgba(0,0,0,0.1);}table{width:100%;border-collapse:collapse;}td,th{padding:12px;text-align:left;border-bottom:1px solid #eee;}a{text-decoration:none;color:#007bff;}.btn{background:#007bff;color:#fff;padding:8px 16px;border-radius:6px;float:right;}</style></head><body><div class='box'>")

	if canUpload {
		fmt.Fprint(w, "<a href='/upload' class='btn'>+ Upload</a>")
	}
	fmt.Fprintf(w, "<h2>Index of %s</h2><table>", urlPath)

	if urlPath != "/" {
		fmt.Fprintf(w, "<tr><td><a href='..'>📁 .. (Parent)</a></td></tr>")
	}

	for _, f := range files {
		icon := "📄 "
		if f.IsDir() {
			icon = "📁 "
		}
		fmt.Fprintf(w, "<tr><td>%s <a href='%s'>%s</a></td></tr>", icon, filepath.Join(urlPath, f.Name()), f.Name())
	}
	fmt.Fprint(w, "</table></div></body></html>")
}
