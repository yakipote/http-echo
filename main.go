package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// パスとハンドラー関数を結びつける
	http.HandleFunc("/", handler())

	// Web サーバーを起動する
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "host:")
		fmt.Fprintf(writer, "- %v/=\n", request.Host)
		fmt.Fprintln(writer, "headers:")
		for i, v := range request.Header {
			fmt.Fprintf(writer, "- %v:%v\n", i, v)
		}
		fmt.Fprintln(writer, "body:")
		_, err := io.Copy(writer, request.Body)
		if err != nil {
			log.Printf("faild to write data")
		}
	}
}
