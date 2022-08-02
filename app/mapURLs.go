package app

import (
	orderService "github.com/greenbahar/manage-order-process/order-service"
)

func mapUrls(httpHandle orderService.Handler) {
	router.POST("api/order", httpHandle.SaveOrder)
}
