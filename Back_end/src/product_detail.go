package src

import (
	models "PJ/SmartDevice_Ecomerce/Back_end/Models"
	repositories "PJ/SmartDevice_Ecomerce/Back_end/Repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SeeProductDetails(c *gin.Context) {
	ProductID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.HTML(400, "notice.html", gin.H{
			"notice": "Fail to find product",
		})
		return
	}
	var productDetail models.ProductDetail

	err = repositories.Instance.
		Table("products").
		Joins("LEFT JOIN reviews on products.product_id = reviews.product_id").
		Select("products.product_id,products.name,products.description,products.price,products.stock_quantity,products.category_id,reviews.rating").
		Where("products.product_id= ?", ProductID).
		Scan(&productDetail).Error

	if err != nil {
		c.HTML(400, "notice.html", gin.H{
			"notice": "Fail to fetching data",
		})
		return
	}

	err = repositories.Instance.
		Table("reviews").
		Select("review_text").
		Where("product_id = ?", ProductID).
		Scan(&productDetail.ReviewText).
		Error
	if err != nil {
		c.HTML(400, "notice.html", gin.H{
			"notice": "Fail to fetching data",
		})
		return
	}

	var categories []models.Category

	categories, err = GetAllCategory()

	if err != nil {
		c.HTML(400, "notice.html", gin.H{
			"notice": "fail to fetching category data",
		})
		return
	}

	name, role, IsAuthenticated := GetSessionData(c)

	c.HTML(200, "product_detail.html", gin.H{
		"name":            name,
		"role":            role,
		"isauthenticated": IsAuthenticated,
		"category":        categories,
		"detail":          productDetail,
	})

}
