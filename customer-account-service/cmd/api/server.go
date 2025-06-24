package main

import (
	accountRepository "customer-account-service/customer-account-service/internal/account/infrastructure/repository"
	accountHandlers "customer-account-service/customer-account-service/internal/account/interfaces/handlers"
	orderRepository "customer-account-service/customer-account-service/internal/orders/infrastructure/repository"
	orderHandlers "customer-account-service/customer-account-service/internal/orders/interfaces/handlers"
	"customer-account-service/customer-account-service/internal/shared/db"
	handlers "customer-account-service/customer-account-service/internal/shared/handler"
	"customer-account-service/customer-account-service/internal/shared/router"
	"fmt"
	"log"
	"os"
	"time"
)

// @title Customer Service API
// @version 1.0
// @description API for managing customer account.

// @schemes https
// @host customer-account-service.onrender.com
// @BasePath /api
func main() {

	dbUsed, err := db.ConnectDb()
	if err != nil {
		log.Fatalln(err)
	}

	accountRepo := accountRepository.New(dbUsed)
	orderRepo := orderRepository.New(dbUsed)

	sharedHandler := &handlers.Handler{
		AccountRepo: accountRepo,
		OrderRepo:   orderRepo,
		Db:          dbUsed,
	}

	accountHandler := &accountHandlers.AccountHandler{
		Handler: sharedHandler,
	}

	orderHandler := &orderHandlers.OrderHandler{
		Handler: sharedHandler,
	}

	mainRouter := router.SetUpRouter(accountHandler, orderHandler)

	port := os.Getenv("API_PORT")

	if port == "" {
		port = "8082"
	}

	fmt.Println(time.Now())

	if err := mainRouter.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
