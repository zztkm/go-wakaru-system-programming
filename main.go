package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type PrintlnWriter struct {
	w io.Writer
}

func (w PrintlnWriter) Write(p []byte) (int, error) {
	return fmt.Fprintln(w.w, string(p))
}


func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json 化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}
	// ここにコードを書く
	respons, err := json.Marshal(source)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "失敗！")
	}
	gw := gzip.NewWriter(w)
	pw := PrintlnWriter{w: os.Stdout}
	mw := io.MultiWriter(gw, pw)
	mw.Write(respons)
	gw.Flush()

}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
