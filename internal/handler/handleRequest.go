package handler

import (
	"encoding/json"
	"httproxy/internal/domain"
	"io"
	"net/http"
	"sync"
	"time"
)

var requestResponseStore sync.Map

// HandleRequest processes incoming proxy requests
// @Summary      Proxy HTTP requests
// @Description  This endpoint handles HTTP requests from clients
// @Tags         requests
// @Accept       json
// @Produce      json
// @Param		 input body domain.ProxyRequest true "proxy request"
// @Success      200  {object}  domain.ProxyResponse
// @Failure      400  {string} string  "Bad Request"
// @Failure      404  {string} string  "Status Not Found"
// @Failure      500  {string} string  "Internal Server Error"
// @Router       /proxy [post]
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	var proxyRequestData domain.ProxyRequest

	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

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
	proxyResponseData := domain.ProxyResponse{
		ID:      requestId,
		Status:  resp.StatusCode,
		Headers: resp.Header,
		Length:  contentLength,
	}

	requestResponseStore.Store(requestId, domain.StoredRequestResponse{
		Request:  proxyRequestData,
		Response: proxyResponseData,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proxyResponseData)
}
