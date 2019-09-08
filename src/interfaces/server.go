package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-rti-testing/src/domain"
	"go-rti-testing/src/infrastructure"
)

func InitServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.HandleFunc("/calculate_offer", func(w http.ResponseWriter, r *http.Request) {
		paramIn := []domain.Condition{}
		err := json.NewDecoder(r.Body).Decode(&paramIn)
		if err != nil {
			infrastructure.Log("error in decode json", err)
		}
		product := getProduct()
		offer, err := domain.Calculate(product, paramIn)
		if err != nil {
			infrastructure.Log("error in make offer", err)
		}
		w.Header().Add("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(offer)
		if err != nil {
			infrastructure.Log("error in marshal json", err)
		}
	})
	log.Printf("server starts")
	err := http.ListenAndServe(":8099", nil)
	if err != nil {
		infrastructure.Log("error in IntServer", err)
	}
}
