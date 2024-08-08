package handlers

import (
	models "PJ/SmartDevice_Ecomerce/Back_end/Models"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Retrieve the product from the database
	cart := models.GetCartFromContext(c)

	cart.AddItem(productID, 1)

	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
}
func RemoveFromCart(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	session := sessions.Default(c)
	sessionID := session.Get("ID")
	cart := models.GetCartFromContext(c)

	for _, item := range cart.Items {
		if item.Product.ProductID == productID {
			quantity := 0 - item.Quantity
			cart.AddItem(productID, quantity)
		}
	}

	models.UpdateCart(sessionID, cart)

	c.Redirect(http.StatusFound, "/cart")
}
