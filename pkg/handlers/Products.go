package handlers

import (
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
	var products []models.Product
	if result := h.Db.Find(&products); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result.Error)
		return
	}
	json.NewEncoder(w).Encode(products)
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
	// add product to database using gorm
	err = h.Db.Create(&product).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("created")

}

func (h handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get product id from url
	var_id := mux.Vars(r)["id"]
	// string to int
	id, _ := strconv.Atoi(var_id)
	// find product in mocks
	var product models.Product
	// marshal product to json
	if result := h.Db.First(&product, id); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(result.Error)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)

}
