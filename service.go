package main

import "errors"

type HelloService interface{
	hello (string) (string, error)
}

type helloService struct{}

func (helloService)hello (s string) (string, error){
	response :="Hello World"
	if s =="" {
		return "", errEmpty
	}
	return response, nil
}

var errEmpty = errors.New("empty request")



