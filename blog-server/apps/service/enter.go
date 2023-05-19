package service

import "blog-server/apps/service/image_svc"

type ServiceGroup struct {
	ImageService image_svc.ImageService
}

var ServiceApp = new(ServiceGroup)
