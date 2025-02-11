package src

import (
	models "PJ/SmartDevice_Ecomerce/Back_end/Models"
	repositories "PJ/SmartDevice_Ecomerce/Back_end/Repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowDashboard(c *gin.Context) {

	name, role, IsAuthenticated := GetSessionData(c)
	if name == nil {
		name = "GUEST"
	}

	cat, err := GetAllCategory()

	if err != nil {
		LogErr(c, 401, "fail to fetching data")
	}

	var parent_cat []models.Category

	err = repositories.Instance.
		Table("categories").
		Select("*").
		Where("parent_category_id IS NULL").Scan(&parent_cat).Error

	if err != nil {
		LogErr(c, 401, "Fail to fetching data")
	}

	c.HTML(200, "index.html", gin.H{
		"name":      name,
		"role":      role,
		"isauth":    IsAuthenticated,
		"cat":       cat,
		"parentCat": parent_cat,
	})
}
func ShowProductOrderbyCategory(c *gin.Context) {
	name, role, IsAuthenticated := GetSessionData(c)
	CtgrId, err := strconv.Atoi(c.Param("ctgrid"))

	if err != nil {
		c.HTML(401, "notice.html", gin.H{
			"notice": "Can't find category",
		})
		return
	}
	var data []models.Product

	err = repositories.Instance.
		Table("products as p").
		Joins("JOIN categories as c ON p.category_id = c.category_id").
		Select("p.product_id,p.name,p.description,p.price,p.stock_quantity").
		Where("p.category_id = ? or c.parent_category_id = ?", CtgrId, CtgrId).
		Scan(&data).
		Order("p.product_id ASC").
		Error
	if err != nil {
		c.HTML(401, "notice.html", gin.H{
			"notice": "Fail to fetching data",
		})
		return
	}

	for i := range data {
		result, err := GetIMG(data[i].ProductID)
		if err != nil {
			LogErr(c, 401, err.Error())
		}
		for j := range result {
			if result[j].IsCover == 0 {
				data[i].IMGpath = append(data[i].IMGpath, result[j].IMGPath)
			} else {
				data[i].Coverpath = result[j].IMGPath
			}
		}

	}

	var category []models.Category
	category, err = GetAllCategory()

	var categoryName string

	for i := range category {
		if category[i].CategoryID == CtgrId {
			categoryName = category[i].CategoryName
		}
	}

	if err != nil {
		c.HTML(401, "notice.html", gin.H{
			"notice": "Fail to fetching data",
		})
		return
	}

	c.HTML(200, "ShowProduct.html", gin.H{
		"products":        data,
		"category":        category,
		"categoryname":    categoryName,
		"name":            name,
		"role":            role,
		"isauthenticated": IsAuthenticated,
	})
}
