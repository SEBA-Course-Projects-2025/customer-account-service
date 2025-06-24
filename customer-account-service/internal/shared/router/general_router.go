package router

import (
	accountInterfaces "customer-account-service/customer-account-service/internal/account/interfaces"
	accountHandlers "customer-account-service/customer-account-service/internal/account/interfaces/handlers"
	orderInterfaces "customer-account-service/customer-account-service/internal/orders/interfaces"
	orderHandlers "customer-account-service/customer-account-service/internal/orders/interfaces/handlers"
	"customer-account-service/customer-account-service/internal/shared/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRouter(accountHandler *accountHandlers.AccountHandler, orderHandler *orderHandlers.OrderHandler) *gin.Engine {

	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("auth/login", accountHandler.LoginAccountHandler)

	api := r.Group("/api", middlewares.AuthMiddleware())

	accountInterfaces.SetUpAccountsRouter(api, accountHandler)
	orderInterfaces.SetUpOrdersRouter(api, orderHandler)

	return r
}
