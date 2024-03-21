package main

import (
	"log"

	"github.com/ashkarax/go-grpc-apigateway/pkg/auth"
	"github.com/ashkarax/go-grpc-apigateway/pkg/config"
	"github.com/ashkarax/go-grpc-apigateway/pkg/order"
	"github.com/ashkarax/go-grpc-apigateway/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at getting config", err)
	}

	r := gin.Default()

	authSvc := auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, authSvc)
	order.RegisterRoutes(r, &c, authSvc)

	r.Run(c.Port)
}
