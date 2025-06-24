package interfaces

import (
	"customer-account-service/customer-account-service/internal/orders/interfaces/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpOrdersRouter(rg *gin.RouterGroup, h *handlers.OrderHandler) {
	orders := rg.Group("orders")
	{
		orders.GET("/", h.GetOrdersHandler)

		orders.GET("/:orderId", h.GetOrderByIdHandler)
		orders.PUT("/:orderId", h.PutOrderStatusHandler)
	}
}
