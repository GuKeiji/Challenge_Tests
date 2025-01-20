package repository

import "app/internal"

func NewProductsMapMock() *ProductsMap {
	mockDB := map[int]internal.Product{
		1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Test id 1", Price: 20, SellerId: 1}},
		2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "Test id 2", Price: 30, SellerId: 2}},
		3: {Id: 3, ProductAttributes: internal.ProductAttributes{Description: "Test id 3", Price: 40, SellerId: 3}},
	}

	return &ProductsMap{
		db: mockDB,
	}
}

type ProductsMapMock struct {
	db map[int]internal.Product
}

func (r *ProductsMapMock) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	p = make(map[int]internal.Product)

	// search the products
	for k, v := range r.db {
		if query.Id > 0 && query.Id != v.Id {
			continue
		}

		p[k] = v
	}

	return
}
