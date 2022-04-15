package handlers

import (
	"diy/pkg/mocks"
	"diy/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (h handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Products)
}

//add procuts to mocks json
func (h handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	// read request body
	var product models.Product
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// unmarshal body
	err = json.Unmarshal(body, &product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// add product to mocks
	mocks.Products = append(mocks.Products, product)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("created")

}

func (h handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	// get product id from url
	var_id := mux.Vars(r)["id"]
	// string to int
	id, _ := strconv.Atoi(var_id)
	// find product in mocks
	for _, product := range mocks.Products {
		if product.Id == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(product)
			return
		}
	}

}
