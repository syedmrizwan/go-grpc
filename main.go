package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/syedmrizwan/go-grpc/env"
	"github.com/syedmrizwan/go-grpc/utils"
)

func main() {
	logger := utils.GetLogger()
	router, err := Inject()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	router.Use(cors.Default())
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))
	if err := router.Run(":" + env.Env.ServerListenPort); err != nil {
		panic(err.Error())
	}
}
