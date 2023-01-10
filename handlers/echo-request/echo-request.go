package echo_request

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Handler() http.HandlerFunc {
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
