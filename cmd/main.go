package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	foodwheel "github.com/tydanny/foodwheel/pkg"
)

func main() {
	router := gin.Default()

	foodwheel.InitializeRoutes(router)

	if err := router.Run(":8080"); err != nil {
		// TODO: don't just panic here. need some sort of logging
		panic(fmt.Errorf("server exited unexpectedly: %v", err))
	}
}
