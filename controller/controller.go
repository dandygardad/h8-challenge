package controller

import "challenge10/service"

type Controller struct {
	service service.ServiceInterface
}

func NewController(service service.ServiceInterface) *Controller {
	return &Controller{service: service}
}
