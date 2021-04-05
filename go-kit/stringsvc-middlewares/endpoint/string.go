package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"stringsvc-middlewares/service"
)

type UppercaseRequest struct {
	S string `json:"s"`
}

type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type CountRequest struct {
	S string `json:"s"`
}

type CountResponse struct {
	V int `json:"v"`
}

func MakeUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return UppercaseResponse{v, err.Error()}, nil
		}
		return UppercaseResponse{v, ""}, nil
	}
}

func MakeCountEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v := svc.Count(req.S)
		return CountResponse{v}, nil
	}
}
