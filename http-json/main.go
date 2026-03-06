package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Payment struct {
	Description string `json:"description"`
	USD         int    `json:"usd"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	Time        time.Time
}

type HttpResponse struct {
	Money          int       `json:"money"`
	PaymentHistory []Payment `json:"paymentHistory"`
}

func (p Payment) Println() {
	fmt.Println("Description: ", p.Description)
	fmt.Println("USD: ", p.USD)
	fmt.Println("FullName: ", p.FullName)
	fmt.Println("Address: ", p.Address)
}

var mtx = sync.Mutex{}
var money = 1000
var paymentHistory = make([]Payment, 0)

func payHandler(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		fmt.Println("Decoder problem: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	payment.Time = time.Now()

	mtx.Lock()
	if money >= payment.USD {
		money -= payment.USD
		paymentHistory = append(paymentHistory, payment)
	}
	mtx.Unlock()
	httpResponse := HttpResponse{
		Money:          money,
		PaymentHistory: paymentHistory,
	}

	b, err := json.Marshal(httpResponse)
	//b, err := json.MarshalIndent(httpResponse, "", "	")

	if err != nil {
		fmt.Println("Marshal problem")
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("Write problem")
	}
}

func main() {
	http.HandleFunc("/pay", payHandler)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Listen error: ", err)
	}
}
