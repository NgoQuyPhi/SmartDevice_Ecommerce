package routers

import (
	middleware "PJ/SmartDevice_Ecomerce/Back_end/Middleware"
	"PJ/SmartDevice_Ecomerce/Back_end/src"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/Front_end", "./Front_end")
	router.LoadHTMLGlob("Front_end/templates/*")
	router.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("secret"))))
	router.GET("/", src.ShowDashboard)
	Cart := router.Group("/cart")
	{
		Cart.Static("/Front_end", "./Front_end")
		Cart.GET("/", src.GetCart)
		Cart.POST("/:id", src.AddToCart)
		Cart.POST("/add/:id", src.Add_item_quantity)
		Cart.POST("/subtract/:id", src.Subtract_item_quantity)
		Cart.POST("/remove/:id", src.RemoveFromCart)
		Cart.POST("/deleteall", src.DeleteAllItemInCart)
	}
	Cate := router.Group("/ctgr")
	{
		Cate.Static("/Front_end", "./Front_end")
		Cate.GET("/:ctgrid", src.ShowProductOrderbyCategory)
	}
	router.GET("/login", src.ShowLoginPage)
	router.POST("/login", middleware.LoginHandle)
	router.GET("/logout", src.Logout)

	router.GET("/signup", src.ShowSignupPage)
	router.POST("/signup", src.SignUp)
	router.GET("/detail/:id", src.SeeProductDetails)
	return router
}
