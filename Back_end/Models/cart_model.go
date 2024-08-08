// data/cart_data.go
package models

import (
	database "PJ/SmartDevice_Ecomerce/Back_end/Database"
	"sync"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	cartData   = make(map[interface{}]*Cart)
	cartDataMu sync.Mutex
)

type Cart struct {
	Items []CartItem
	Total float32
}

type CartItem struct {
	Product    Product
	CoverPhoto string
	Quantity   int
}

// AddItem adds a product to the cart or updates the quantity if it already exists
func (c *Cart) AddItem(productID int, quantity int) {
	for i, item := range c.Items {
		if item.Product.ProductID == productID {
			c.Items[i].Quantity += quantity
			c.Total += item.Product.Price * float32(quantity)
			if c.Items[i].Quantity == 0 {
				var newItems []CartItem
				for _, item := range c.Items {
					if item.Product.ProductID != productID {
						newItems = append(newItems, item)
					}
				}

				c.Items = newItems
			}
			return
		}
	}
	var product Product
	err := database.Instance.
		Table("products").
		Select("product_id", "name", "price", "description").
		Where("product_id = ?", productID).
		Scan(&product).Error
	if err != nil {
		return
	}
	c.Items = append(c.Items, CartItem{Product: product, Quantity: 1})
	c.Total = c.Total + product.Price
}

// GetCart retrieves the cart data for a given session ID
func GetCart(sessionID interface{}) *Cart {
	cartDataMu.Lock()
	defer cartDataMu.Unlock()
	cart, ok := cartData[sessionID]
	if !ok {
		cart = &Cart{
			Items: []CartItem{},
		}
		cartData[sessionID] = cart
	}
	return cart
}

func UpdateCart(sessionID interface{}, cart *Cart) {
	cartDataMu.Lock()
	defer cartDataMu.Unlock()
	cartData[sessionID] = cart
}

func GetCartFromContext(c *gin.Context) *Cart {
	// Retrieve the session ID from the request context
	sessionID := getSessionID(c)

	// Get the cart data for the current session
	return GetCart(sessionID)
}

// getSessionID retrieves the session ID from the request context
// You'll need to implement this function based on your session management approach
func getSessionID(c *gin.Context) interface{} {
	session := sessions.Default(c)
	sessionID := session.Get("ID")
	if sessionID == nil {
		sessionID = ""
	}
	return sessionID

}
