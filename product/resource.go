package product

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type Resource struct {
	Db *gorm.DB
}

func (resource *Resource) FindAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed."))
		return
	}

	var resources []Product
	resource.Db.Preload("Reviews").Find(&resources)

	response := make(map[string]interface{})
	response["products"] = resources

	b, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "INTERNAL_ERROR", 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write(b)
}

func (resource *Resource) Find(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed."))
		return
	}

	id, _ := strconv.Atoi(r.URL.Path[len("/api/products/"):])

	if id <= 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Content not found."))
		return
	}

	var result Product
	resource.Db.Preload("Reviews").Find(&result, id)

	if result.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Content not found."))
		return
	}

	response := make(map[string]interface{})
	response["product"] = result
	b, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "INTERNAL_ERROR", 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write(b)
}
