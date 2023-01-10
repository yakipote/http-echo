package random

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

const M = 1000 * 1000
const maxSize = 10 * M

func Handler() http.HandlerFunc {
	src := rand.NewSource(0)
	return func(writer http.ResponseWriter, request *http.Request) {
		queries := request.URL.Query()
		size := queries.Get("size")
		max, err := strconv.Atoi(size)
		if err != nil {
			max = maxSize
		}
		read := rand.New(src)
		_, err = io.CopyN(writer, read, int64(max))
		if err != nil {
			log.Printf("faild to write random data %s", err)
			return
		}
	}
}
