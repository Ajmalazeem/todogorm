package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type helloRequest struct{
	S string 
}

type helloResponse struct{
	Display string `json:"display"`
}

func makeHelloEndpoint (svc HelloService) endpoint.Endpoint{
	return func(_ context.Context, request interface{},)(interface{},error){
		req := request.(helloRequest)
		v , err := svc.hello(req.S)
		if err !=nil{
			return helloResponse{err.Error()},nil
		}
	 return helloResponse{v},nil
	}
}

func decodeHelloResponse(_ context.Context,r *http.Request)(interface{},error){
	var request helloRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err !=nil {
		return nil, err
	}
	return request,nil

}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{})error{
	return json.NewEncoder(w).Encode(response)

}