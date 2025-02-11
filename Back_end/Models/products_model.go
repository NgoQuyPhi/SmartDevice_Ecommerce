package models

import "time"

type Product struct {
	ProductID      int         `json:"product_id"`
	Name           string      `json:"name"`
	Price          float32     `json:"price"`
	Description    string      `json:"description"`
	Stock_quantity int         `json:"stock_quantity"`
	CategoryID     int         `json:"category_id"`
	Create_at      *time.Timer `json:"create_at"`
	Update_at      *time.Timer `json:"update_at"`
	IMGpath        []string
	Coverpath      string
}

type ProductIMG struct {
	IMGPath string `json:"image_path"`
	IsCover int    `json:"iscover"`
}

type Reviews struct {
	ReviewID  int `json:"review_id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	ShowReviews
}

type ShowReviews struct {
	Rating     int    `json:"rating"`
	ReviewText string `json:"review_text"`
}
type Category struct {
	CategoryID       int    `json:"category_id"`
	CategoryName     string `json:"category_name"`
	ParentCategoryID int    `json:"parent_category_id"`
	Cover            string `json:"cover"`
}

type ProductDetail struct {
	ProductID      int     `json:"product_id"`
	Name           string  `json:"name"`
	Price          float32 `json:"price"`
	Description    string  `json:"description"`
	Stock_quantity int     `json:"stock_quantity"`
	CategoryID     int     `json:"category_id"`
	Rating         int     `json:"rating"`
	CoverIMG       string
	ReviewText     []string `json:"review_text"`
}
