package handlers

import (
	database "PJ/SmartDevice_Ecomerce/Back_end/Database"
	models "PJ/SmartDevice_Ecomerce/Back_end/Models"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowDashboard(c *gin.Context) {

	session := sessions.Default(c)
	IsAuthenticated := session.Get("isauthenticated")

	name := session.Get("name")
	role := session.Get("role")

	var data []models.Product

	err := database.Instance.
		Table("products").
		Select("product_id,name,description,price,stock_quantity").
		Scan(&data).
		Order("product_id ASC").
		Error
	if err != nil {
		c.HTML(401, "notice.html", gin.H{
			"notice": "Fail to fetching products data",
		})
		return
	}

	var category []models.Category

	err = database.Instance.
		Table("categories").
		Select("*").
		Scan(&category).
		Error

	if err != nil {
		c.HTML(401, "notice.html", gin.H{
			"notice": "fail to fetching category data",
		})
		return
	}

	var coverPhotopath []string

	err = database.Instance.
		Table("product_images").
		Select("image_path").
		Where("IsCover = 1").
		Order("product_id ASC").
		Scan(&coverPhotopath).
		Error
	if err != nil {
		c.HTML(401, "notice.html", gin.H{
			"notice": "Fail to fetching data",
		})
		return
	}
	for i := range data {
		data[i].CoverIMG = coverPhotopath[i]
	}
	c.HTML(200, "dashboard.html", gin.H{
		"products":        data,
		"category":        category,
		"name":            name,
		"role":            role,
		"isauthenticated": IsAuthenticated,
	})
}
func ShowProductOrderbyCategory(c *gin.Context) {
	session := sessions.Default(c)
	IsAuthenticated := session.Get("isauthenticated")
	name := session.Get("name")
	role := session.Get("role")
	CtgrId, err := strconv.Atoi(c.Param("ctgrid"))

	if err != nil {
		c.HTML(401, "notice.html", gin.H{
			"notice": "Can't find category",
		})
		return
	}
	var data []models.Product

	err = database.Instance.
		Table("products").
		Joins("JOIN categories ON products.category_id = categories.category_id").
		Select("product_id,products.name,description,price,stock_quantity").
		Where("products.category_id = ? or parent_category_id = ?", CtgrId, CtgrId).
		Scan(&data).
		Order("product_id ASC").
		Error
	if err != nil {
		c.HTML(401, "notice.html", gin.H{
			"notice": "Fail to fetching data",
		})
		return
	}

	var category []models.Category

	err = database.Instance.
		Table("categories").
		Select("*").
		Scan(&category).
		Error

	if err != nil {
		c.HTML(401, "notice.html", gin.H{
			"notice": "fail to fetching category data",
		})
		return
	}

	var coverPhotopath []string

	err = database.Instance.
		Table("product_images").
		Select("image_path").
		Where("IsCover = 1").
		Order("product_id ASC").
		Scan(&coverPhotopath).
		Error
	if err != nil {
		c.HTML(401, "notice.html", gin.H{
			"notice": "Fail to fetching data",
		})
		return
	}
	for i := range data {
		data[i].CoverIMG = coverPhotopath[i]
	}
	c.HTML(200, "dashboard.html", gin.H{
		"products":        data,
		"category":        category,
		"name":            name,
		"role":            role,
		"isauthenticated": IsAuthenticated,
	})
}
