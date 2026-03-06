package main

import (
	"fmt"
	"net/http"
	"time"
)

func payHandler(w http.ResponseWriter, r *http.Request) {
	//str := []string{"Pay asd", "123", "asdaa"}
	str := "Pay asd"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("error writing response ", err.Error())
	} else {
		fmt.Println("Pay response written")
	}
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "Cancel"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("error writing response ", err.Error())
	} else {
		fmt.Println("Cancel response written")
	}
}
func check(w http.ResponseWriter, request *http.Request) {
	_, err := w.Write([]byte("Hello Worldaaa"))
	if err != nil {
		fmt.Println("error writing response ", err.Error())
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	str := "Hello world"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("error writing response ", err.Error())
	} else {
		fmt.Println("response written")
	}
}

func handlerSleep(w http.ResponseWriter, request *http.Request) {
	time.Sleep(5 * time.Second)
	str := "Hello world"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("error writing response ", err.Error())
	} else {
		fmt.Println("response written")
	}
}

func main() {
	http.HandleFunc("/", check)
	http.HandleFunc("/default", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)
	http.HandleFunc("/sleep", handlerSleep)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("error starting server", err.Error())
	}
}
