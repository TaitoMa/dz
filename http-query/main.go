package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fooParam := r.URL.Query().Get("foo")
	booParam := r.URL.Query().Get("boo")

	fmt.Println("foo: ", fooParam, "boo: ", booParam)
}

func main() {
	http.HandleFunc("/default", defaultHandler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Listen error: ", err)
	}
}
