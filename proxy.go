package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type ProxyRequest struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type ProxyResponse struct {
	ID      string              `json:"id"`
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers"`
	Length  int64               `json:"length"`
}

type StoredRequestResponse struct {
	Request  ProxyRequest
	Response ProxyResponse
}

var requestResponseStore sync.Map

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var proxyRequestData ProxyRequest
	err := json.NewDecoder(r.Body).Decode(&proxyRequestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	proxyReq, err := http.NewRequest(proxyRequestData.Method, proxyRequestData.URL, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for key, value := range proxyRequestData.Headers {
		proxyReq.Header.Set(key, value)
	}
	customTransport := &http.Transport{}

	resp, err := customTransport.RoundTrip(proxyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	contentLength := int64(len(response))

	requestId := time.Now().Format("20060102150405")
	proxyResponseData := ProxyResponse{
		ID:      requestId,
		Status:  resp.StatusCode,
		Headers: resp.Header,
		Length:  contentLength,
	}

	requestResponseStore.Store(requestId, StoredRequestResponse{
		Request:  proxyRequestData,
		Response: proxyResponseData,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proxyResponseData)
}

func getStoredRequestHandler(w http.ResponseWriter, r *http.Request) {
	var storedRequest []StoredRequestResponse

	requestResponseStore.Range(func(key, value interface{}) bool {
		storedRequest = append(storedRequest, value.(StoredRequestResponse))
		return true
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storedRequest)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/proxy", handleRequest)
	mux.HandleFunc("/stored_requests", getStoredRequestHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Listening on :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
