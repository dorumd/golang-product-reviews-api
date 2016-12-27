package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// ApiStatusHandler - /api/status
func ApiStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed."))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

// GetProductsHandler - /api/products?page=1
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed."))
		return
	}

	limit := 10
	// offset := 0

	// Get Limit & Offset
	limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
	// offset, _ = strconv.Atoi(r.URL.Query().Get("offset"))

	response := make(map[string]interface{})
	response["products"] = make([]Product, limit)

	b, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "INTERNAL_ERROR", 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write(b)
}

// GetProductReviewsHandler - /api/product-reviews/1
func GetProductReviewsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed."))
		return
	}

	id, _ := strconv.Atoi(r.URL.Path[len("/api/product-reviews/"):])

	if id <= 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Content not found."))
		return
	}

	response := make(map[string]interface{})
	response["product"] = Product{ Id: id }

	b, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "INTERNAL_ERROR", 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	w.Write(b)
}
