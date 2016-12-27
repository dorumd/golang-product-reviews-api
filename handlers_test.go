package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
)

func TestApiStatusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ApiStatusHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestApiStatusHandlerWithWrongMethod(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ApiStatusHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}


	// Check the response body is what we expect.
	expected := "Method not allowed."
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetProductsHandler(t *testing.T) {
	limit := 10
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/products?limit=%v", limit), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProductsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := make(map[string]interface{})
	expected["products"] = make([]Product, limit)
	test, _ := json.Marshal(expected)

	if rr.Body.String() != string(test) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), string(test))
	}
}

func TestGetProductsHandlerWithWrongMethod(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProductsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}


	// Check the response body is what we expect.
	expected := "Method not allowed."
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetProductReviewsHandler(t *testing.T) {
	id := 1
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/product-reviews/%v", id), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProductReviewsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	response := make(map[string]interface{})
	response["product"] = Product{ Id: id }
	expected, _ := json.Marshal(response)

	if rr.Body.String() != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
	}
}

func TestGetProductReviewsHandlerWithWrongMethod(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/product-reviews/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProductReviewsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}


	// Check the response body is what we expect.
	expected := "Method not allowed."
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}


func TestGetProductReviewsHandlerWithoutId(t *testing.T) {
	req, err := http.NewRequest("GET","/api/product-reviews/",nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProductReviewsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}