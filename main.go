package main

import (
	"github.com/gin-gonic/gin"
	api "github.com/tydanny/foodwheel/v1alpha1"
)

func main() {
	router := gin.Default()

	api.InitializeRoutes(router)

	router.Run(":3000")
}
