package model

type Product struct {
	Product_Id string `json:"product_id"`
	Product_Name string  `json:"product_name"`
	Product_Brand string `json:"product_brand"`
	Product_Img string `json:"product_img"`
	Product_Price string `json:"product_price"`
}

func (pr Product)GetBrand() string{
	return pr.Product_Brand
}