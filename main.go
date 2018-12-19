package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httputil"

	"./iris"
)

var orders = make(map[string]iris.Order)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/orders", addOrder).Methods(http.MethodPost)
	router.HandleFunc("/orders", getOrders).Methods(http.MethodGet)
	router.HandleFunc("/orders/{orderId}", getOrder).Methods(http.MethodGet)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Listening...")

	server.ListenAndServe()
}

func getOrder(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	orderId := params["orderId"]
	if order, ok := orders[orderId]; ok {
		json.NewEncoder(writer).Encode(order)
	} else {
		http.Error(writer, "Order not found for given orderId", http.StatusNotFound)
	}

}

func getOrders(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(orders)
}

func addOrder(writer http.ResponseWriter, request *http.Request) {
	var order iris.Order

	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		http.Error(writer, "Can't dump request", http.StatusInternalServerError)
	} else {
		log.Printf("Inbound Request:\n%s", string(requestDump))
	}

	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&order)
	if err != nil {
		log.Printf("Error: %s", err)
		http.Error(writer, "Can't decode OrderDTO", http.StatusBadRequest)
	}

	log.Printf("Adding order: %s", order.OrderId)
	orders[order.OrderId] = order

	writer.WriteHeader(http.StatusCreated)
}
