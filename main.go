package main

import (
	echo_request "http-echo/handlers/echo-request"
	"http-echo/handlers/random"
	"log"
	"net/http"
)

func main() {
	// パスとハンドラー関数を結びつける
	http.HandleFunc("/", echo_request.Handler())
	http.HandleFunc("/echo", echo_request.Handler())
	http.HandleFunc("/random", random.Handler())

	// Web サーバーを起動する
	log.Fatal(http.ListenAndServe(":8080", nil))
}
