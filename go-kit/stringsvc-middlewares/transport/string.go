package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"stringsvc-middlewares/endpoint"
)

func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
