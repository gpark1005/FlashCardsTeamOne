package handler


import (

)

type Service interface {

}

type InfoHandler struct {
	Svc Service
}

func NewInfoHandler(s Service) InfoHandler{
	return InfoHandler{
		Svc: s,
	}
}