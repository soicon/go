package service

import (
	"github.com/dmr/microservice/productservice/dbclient"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"github.com/dmr/microservice/productservice/model"
	"fmt"
)

var DBClient dbclient.IBoltClient
func GetBrand(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Starting querying Brand")
	// Read the account struct BoltDB
	account, err := DBClient.QueryBrand()

	// If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(account)
	fmt.Print(data)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {

	// Read the 'accountId' path parameter from the mux map
	brand := mux.Vars(r)["brand"]

	// Read the account struct BoltDB
	account, err := DBClient.QueryAllProduct(brand)

	// If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {

	// Read the 'accountId' path parameter from the mux map
	var accountId = mux.Vars(r)["productId"]

	// Read the account struct BoltDB
	account, err := DBClient.QueryProduct(accountId)

	// If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {

	// Read the 'accountId' path parameter from the mux map
	product := model.Product{}
	json.NewDecoder(r.Body).Decode(&product)
	// Read the account struct BoltDB
	account, err := DBClient.NewProduct(product)

	// If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}