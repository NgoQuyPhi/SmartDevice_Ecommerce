package handlers

import (
	database "PJ/SmartDevice_Ecomerce/Back_end/Database"
	models "PJ/SmartDevice_Ecomerce/Back_end/Models"

	"github.com/gin-gonic/gin"
)

func ShowDashboard(c *gin.Context) {
	var data []models.Product

	err := database.Instance.
		Table("products").
		Select("product_id,name,description,price,stock_quantity,category_id").
		Scan(&data).
		Order("product_id ASC").
		Error
	if err != nil {
		c.HTML(401, "notice.html", gin.H{
			"notice": "Fail to fetching data",
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
		"products": data,
	})
}
