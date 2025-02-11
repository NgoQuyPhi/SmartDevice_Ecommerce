package src

import (
	models "PJ/SmartDevice_Ecomerce/Back_end/Models"
	repositories "PJ/SmartDevice_Ecomerce/Back_end/Repositories"
)

func GetIMG(id int) ([]models.ProductIMG, error) {
	var IMG []models.ProductIMG
	err := repositories.Instance.
		Table("product_images").
		Select("image_path,iscover").
		Where("product_id = ?", id).
		Scan(&IMG).
		Error

	if err != nil {
		return nil, err
	}
	return IMG, nil
}
