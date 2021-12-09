package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func main() {
	reqBody := bytes.NewBufferString("abcdefghijklmnopqrstuvwxyz")
	req := httptest.NewRequest(http.MethodGet, "http://dummy.url.com/", reqBody)
	got := httptest.NewRecorder()

	myHandler(got, req)

	fmt.Println(got.Body.String())
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	body := http.MaxBytesReader(w, r.Body, 10)
	i, err := io.Copy(w, body)
	fmt.Println(i)
	if err != nil {
		fmt.Println(err)
	}
}
