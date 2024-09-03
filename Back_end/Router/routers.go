package routers

import (
	handlers "PJ/SmartDevice_Ecomerce/Back_end/Handlers"
	middleware "PJ/SmartDevice_Ecomerce/Back_end/Middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("Front_end/*")
	router.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("secret"))))
	router.GET("/", handlers.ShowDashboard)
	router.GET("/cart", handlers.GetCart)
	router.POST("cart/:id", handlers.AddToCart)
	router.POST("cart/add/:id", handlers.Add_item_quantity)
	router.POST("cart/subtract/:id", handlers.Subtract_item_quantity)
	router.POST("/cart/remove/:id", handlers.RemoveFromCart)
	router.GET("/login", handlers.ShowLoginPage)
	router.POST("/login", middleware.LoginHandle)
	router.GET("/logout", handlers.Logout)
	router.GET("/ctgr/:ctgrid", handlers.ShowProductOrderbyCategory)
	return router
}
