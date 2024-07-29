package handler

import (
	"encoding/json"
	"httproxy/internal/domain"
	"net/http"
)

// GetStoredRequestHandler returns the stored proxy requests and responses
// @Summary      Get Stored Requests
// @Description  Returns a list of stored proxy requests and responses
// @Tags         requests
// @Produce      json
// @Success      200  {array}   domain.StoredRequestResponse
// @Router       /stored [get]
func GetStoredRequestHandler(w http.ResponseWriter, r *http.Request) {
	var storedRequest []domain.StoredRequestResponse

	requestResponseStore.Range(func(key, value interface{}) bool {
		storedRequest = append(storedRequest, value.(domain.StoredRequestResponse))
		return true
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storedRequest)
}
