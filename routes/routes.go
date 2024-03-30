package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hardikko/ecommerce/controllers"
)

func UserRoutes(inRoutes *gin.Engine) {
	inRoutes.POST("/users/signup", controllers.Signup())
	inRoutes.POST("/users/login", controllers.Login())
	inRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	inRoutes.GET("/users/productview", controllers.SearchProduct())
	inRoutes.GET("/users/search", controllers.SearchProductByQuery())

}
