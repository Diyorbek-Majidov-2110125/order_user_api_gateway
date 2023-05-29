package api

import (
	"practice1/order_user_api_gateway/api/handlers"
	"practice1/order_user_api_gateway/config"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(h handlers.Handler, cfg config.Config) (r *gin.Engine) {
	r = gin.New()

	r.Use(gin.Logger(), gin.Recovery())
//Order
	r.POST("/createOrder",h.CreateOrder)
	r.GET("/getOrder",h.GetOrderById)
	r.GET("/getAllOrder",h.GetAllOrders)
	r.DELETE("/deleteOrder",h.DeleteOrder)


//User
	r.POST("/createUser",h.CreateUser)
	r.GET("/getUser",h.GetUserById)
	r.GET("/getAllUser",h.GetAllUsers)
	r.DELETE("/deleteUser",h.DeleteUser)

//Product
	r.POST("/createProduct",h.CreateProduct)
	r.GET("/getProduct",h.GetProductById)
	r.GET("/getAllProduct",h.GetAllProducts)
	r.DELETE("/deleteProduct",h.DeleteProduct)
	
	
	return
}