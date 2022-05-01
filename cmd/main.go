package main

import (
	"github.com/gin-gonic/gin"
	foodwheel "github.com/tydanny/foodwheel/pkg"
)

func main() {
	router := gin.Default()

	foodwheel.InitializeRoutes(router)

	router.Run(":3000")
}
