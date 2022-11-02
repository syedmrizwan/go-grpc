package main

import (
	"github.com/gin-gonic/gin"
	httpdelivery "github.com/syedmrizwan/go-grpc/app/delivery/http"
)

func Inject() (*gin.Engine, error) {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	httpdelivery.NewHandler(&httpdelivery.Config{
		R:       router,
		BaseURL: "",
	})

	return router, nil
}
