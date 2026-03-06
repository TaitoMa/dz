package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var mtx = sync.Mutex{}
var money = 1000
var bank = atomic.Int64{}

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	fmt.Println(httpRequestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "error writing response " + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	paymentAmount, _ := strconv.Atoi(httpRequestBodyString)
	fmt.Println(paymentAmount)

	mtx.Lock()
	if money >= paymentAmount {
		time.Sleep(3 * time.Second)
		money -= paymentAmount
		fmt.Println("Operation success, money: ", money)
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "error writing response " + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
		}
	}

	httpRequestBodyString := string(httpRequestBody)
	paymentAmount, _ := strconv.Atoi(httpRequestBodyString)

	mtx.Lock()
	if money >= paymentAmount {
		money -= paymentAmount
		bank.Add(int64(paymentAmount))
		fmt.Println("Save success, money: ", bank.Load())
	}
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("error starting server", err.Error())
	}
}
