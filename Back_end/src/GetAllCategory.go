package src

import (
	models "PJ/SmartDevice_Ecomerce/Back_end/Models"
	repositories "PJ/SmartDevice_Ecomerce/Back_end/Repositories"
)

func GetAllCategory() ([]models.Category, error) {
	var category []models.Category

	err := repositories.Instance.
		Table("categories").
		Select("*").
		Scan(&category).
		Error

	if err != nil {
		return nil, err
	}

	return category, nil

}
