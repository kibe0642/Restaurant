package routes

import (
	controller "golang-Restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:orderItem_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("orderItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("orderItems/:orderItem_id", controller.UpdateOrderItem())
}
