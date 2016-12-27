package main

import (
	"log"
	"net/http"
)

func init() {

}

const PORT = ":8080"

func main() {
	log.Printf("Running server on %#v", PORT)

	// Register Handlers
	http.HandleFunc("/api/status", ApiStatusHandler)
	http.HandleFunc("/api/products", GetProductsHandler)
	http.HandleFunc("/api/product-reviews/", GetProductReviewsHandler)

	log.Fatal(http.ListenAndServe(PORT, nil))
}
