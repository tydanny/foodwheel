package v1alpha1

import "github.com/gin-gonic/gin"

func InitializeRoutes(router *gin.Engine) {
	router.GET("/cuisines", getCuisines)
	router.GET("/cuisines/:name", getCuisineByName)
	router.POST("/cuisines", postCuisines)

	router.GET("/spin", getSpin)
}
