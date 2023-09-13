package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"//获取系统变量才用到
	"path/filepath"
)

func main() {
	http.HandleFunc("/", fileHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	dir := os.Getenv("HOME") + "/downloads"//获取系统目录调试时才用
	// dir := "/downloads"//在容器中调试时目录
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	if len(files) == 0 {
		fmt.Fprint(w, "No files available for download.")
		return
	}

	fmt.Fprintln(w, "<html><h2>Available files for download:</h2>")
	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		fileURL := fmt.Sprintf("http://%s/%s", r.Host, file.Name())
		fmt.Fprintf(w, "<a href=\"%s\">%s</a><br></html>", fileURL, file.Name())

		http.HandleFunc("/"+file.Name(), func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, filePath)
		})
	}
}

