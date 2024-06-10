package services

import (
	"product-search-service/models"
	"sort"
	"strings"
)

var products = []models.Product{
	{ID: "1", Name: "Laptop", Brand: "BrandA", Price: 1000, Category: "Electronics", Rating: 4.5},
	{ID: "2", Name: "Smartphone", Brand: "BrandB", Price: 500, Category: "Electronics", Rating: 4.0},
	{ID: "3", Name: "Headphones", Brand: "BrandC", Price: 150, Category: "Accessories", Rating: 4.2},
	{ID: "4", Name: "Smartwatch", Brand: "BrandD", Price: 200, Category: "Wearables", Rating: 3.9},
	{ID: "5", Name: "Tablet", Brand: "BrandE", Price: 300, Category: "Electronics", Rating: 4.3},
	{ID: "6", Name: "Camera", Brand: "BrandF", Price: 800, Category: "Photography", Rating: 4.7},
	{ID: "7", Name: "Speaker", Brand: "BrandG", Price: 120, Category: "Audio", Rating: 4.1},
	{ID: "8", Name: "Keyboard", Brand: "BrandH", Price: 70, Category: "Accessories", Rating: 4.0},
	{ID: "9", Name: "Mouse", Brand: "BrandI", Price: 40, Category: "Accessories", Rating: 3.8},
	{ID: "10", Name: "Monitor", Brand: "BrandJ", Price: 220, Category: "Electronics", Rating: 4.4},
}

const CacheDuration = 5 * 60 //5 min

func SearchProducts(keyword string, priceMin, priceMax float64, category, sortBy string) []models.Product {
	var filteredProducts []models.Product

	for _, product := range products {
		if (keyword == "" || contains(product.Name, keyword) || contains(product.Brand, keyword)) &&
			(priceMin == 0 || product.Price >= priceMin) &&
			(priceMax == 0 || product.Price <= priceMax) &&
			(category == "" || strings.EqualFold(product.Category, category)) {
			filteredProducts = append(filteredProducts, product)
		}
	}

	if sortBy != "" {
		sortProducts(filteredProducts, sortBy)
	}

	return filteredProducts
}

func contains(source, keyword string) bool {
	return strings.Contains(strings.ToLower(source), strings.ToLower(keyword))
}

func sortProducts(products []models.Product, sortBy string) {
	switch sortBy {
	case "price":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price < products[j].Price
		})
	case "rating":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Rating > products[j].Rating
		})
		//adicionar mais classificação se precisar
	}
}
