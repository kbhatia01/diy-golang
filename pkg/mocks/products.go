package mocks

import "diy/pkg/models"

var Products = []models.Product{
	{
		Id:          1,
		Name:        "Product 1",
		Description: "Description 1",
		Price:       1.0,
		Quantity:    1,
	},
	{
		Id:          2,
		Name:        "Product 2",
		Description: "Description 2",
		Price:       2.0,
		Quantity:    2,
	},
	{
		Id:          3,
		Name:        "Product 3",
		Description: "Description 3",
		Price:       3.0,
		Quantity:    3,
	},
}
