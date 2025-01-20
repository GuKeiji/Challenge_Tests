package repository

import (
	"app/internal"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductsMap_SearchProducts(t *testing.T) {
	mockDB := map[int]internal.Product{
		1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Test id 1", Price: 20, SellerId: 1}},
		2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "Test id 2", Price: 30, SellerId: 2}},
		3: {Id: 3, ProductAttributes: internal.ProductAttributes{Description: "Test id 3", Price: 40, SellerId: 3}},
	}

	productsMap := NewProductsMap(mockDB)

	t.Run("Get One Test success", func(t *testing.T) {
		query := internal.ProductQuery{Id: 1}
		product, _ := productsMap.SearchProducts(query)
		expectedProduct := internal.Product{Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Test id 1", Price: 20, SellerId: 1}}

		require.Equal(t, expectedProduct, product[1])
	})

	t.Run("Get One Test failed", func(t *testing.T) {

		query := internal.ProductQuery{Id: 999}
		product, _ := productsMap.SearchProducts(query)

		require.Equal(t, 0, len(product))
	})

	t.Run("Get All Test", func(t *testing.T) {

		query := internal.ProductQuery{}
		product, _ := productsMap.SearchProducts(query)

		for i := range product {
			require.Equal(t, mockDB[i], product[i])
		}
	})
}

func TestNewProductsMap(t *testing.T) {
	mockDB := map[int]internal.Product{
		1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "Test id 1", Price: 20, SellerId: 1}},
		2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "Test id 2", Price: 30, SellerId: 2}},
		3: {Id: 3, ProductAttributes: internal.ProductAttributes{Description: "Test id 3", Price: 40, SellerId: 3}},
	}

	t.Run("TestNewProductsMap sucess", func(t *testing.T) {
		productsMap := NewProductsMap(mockDB)

		require.Equal(t, 3, len(productsMap.db))
	})

	t.Run("TestNewProductsMap failed", func(t *testing.T) {
		productsMap := NewProductsMap(nil)

		require.Equal(t, 0, len(productsMap.db))
	})

}
