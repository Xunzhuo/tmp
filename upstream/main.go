package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var gldata = ""

func store(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	gldata += string(data) + "\n"
	w.Write([]byte("store"))
}

func retrieve(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(gldata))
}

func auth(w http.ResponseWriter, req *http.Request) {
	token := req.Header.Get("token")
	if token != "hello" {
		w.Write([]byte("Unauthorized, token recieved: " + token))
	} else {
		w.Write([]byte("Authorized"))
	}
}
func hello(w http.ResponseWriter, req *http.Request) {
	headers(w, req)
	w.Write([]byte("Hello Envoy!"))
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/store", store)
	http.HandleFunc("/retrieve", retrieve)
	http.HandleFunc("/auth", auth)
	http.HandleFunc("/", hello)
	http.ListenAndServe(":9001", nil)
}
