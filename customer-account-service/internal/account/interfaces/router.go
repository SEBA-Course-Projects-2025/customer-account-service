package interfaces

import (
	"customer-account-service/customer-account-service/internal/account/interfaces/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpAccountsRouter(rg *gin.RouterGroup, h *handlers.AccountHandler) {
	accounts := rg.Group("account")
	{
		accounts.GET("/", h.GetAccountHandler)
		accounts.PUT("/", h.PutAccountHandler)
		accounts.PATCH("/", h.PatchAccountHandler)
	}

}
