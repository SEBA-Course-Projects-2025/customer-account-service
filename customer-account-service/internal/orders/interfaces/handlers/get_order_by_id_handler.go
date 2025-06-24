package handlers

import (
	"customer-account-service/customer-account-service/internal/orders/application/services"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

// GetOrderByIdHandler godoc
// @Summary      Get order by ID
// @Description  Returns a single order by its ID for the given customer.
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer access token"
// @Param        orderId path string true "Order ID (UUID)"
// @Success      200 {object} dtos.OneOrderResponse
// @Failure      400 {object} map[string]interface{} "Invalid customerId or orderId"
// @Failure      404 {object} map[string]interface{} "Order not found"
// @Failure      500 {object} map[string]interface{}
// @Router       /orders/{orderId} [get]
func (h *OrderHandler) GetOrderByIdHandler(c *gin.Context) {

	v, _ := c.Get("customerId")
	customerId, ok := v.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customerId"})
		return
	}

	orderIdStr := c.Param("orderId")

	orderId, err := uuid.Parse(orderIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order UUID"})
		return
	}

	order, err := services.GetOrderById(c.Request.Context(), h.OrderRepo, orderId, customerId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)

}
