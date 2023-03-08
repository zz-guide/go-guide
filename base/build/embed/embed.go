package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed version.txt
var version string

//go:embed version.txt
var version1 []byte

//go:embed version.txt
var embededFiles1 embed.FS

func main() {
	staticServer()
}

func d() {
	fmt.Printf("version: %q\n", version)
	fmt.Printf("version1: %q\n", version1)
	s, _ := embededFiles1.ReadFile("version1.txt")
	fmt.Printf("embededFiles: %s\n", s)
}

func staticServer() {
	// go run . live 磁盘修改文件会影响输出
	// go run . 即使磁盘修改文件也不会影响输出
	useOS := len(os.Args) > 1 && os.Args[1] == "live"
	http.Handle("/", http.FileServer(getFileSystem(useOS)))
	_ = http.ListenAndServe(":9992", nil)
}

//go:embed static
var embededFiles embed.FS

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		log.Print("using live mode")
		return http.FS(os.DirFS("static"))
	}

	log.Print("using embed mode")

	fsys, err := fs.Sub(embededFiles, "static")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}
