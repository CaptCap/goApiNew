package main

import (
	"log"
	"net/http"
)

type wallet struct {
	CreditCards []creditCard
}

type creditCard struct {
	Number         int
	ExpirationDate string
	CvvCode        int
	Holder         string
}

var ivankoCreditCard = creditCard{
	Number:         4224580604397698,
	ExpirationDate: "20 липня 2031р",
	CvvCode:        230,
	Holder:         "Іванко",
}

var ivankoMonoCreditCard = creditCard{
	Number:         4114391864397698,
	ExpirationDate: "20 липня 2025р",
	CvvCode:        892,
	Holder:         "Іванко",
}

var ivankoPrivatCreditCard = creditCard{
	Number:         3911391723597698,
	ExpirationDate: "8 вересня 2025р",
	CvvCode:        777,
	Holder:         "Іванко",
}

var myWallet = wallet{
	CreditCards: []creditCard{
		ivankoMonoCreditCard,
		ivankoPrivatCreditCard,
		ivankoCreditCard,
	},
}

func main() {
	http.HandleFunc("/", api)
	server := &http.Server{
		Addr: ":8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func api(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, my first go api"))
}
