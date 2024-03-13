package main

import (
	"assignment-2/db"
	"assignment-2/handler"
	"assignment-2/repo"
	"assignment-2/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db, err := db.ConnectDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	orderRepo := repo.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)

	orderHandler := handler.NewOrderHandler(orderService)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/orders", orderHandler.CreateOrder)
		v1.GET("/orders", orderHandler.GetOrders)
		v1.PUT("/orders/:id", orderHandler.UpdateOrder)
		v1.DELETE("/orders/:id", orderHandler.DeleteOrder)
	}

	r.Run(":8080")
}
