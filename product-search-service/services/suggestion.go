package services

import (
	"product-search-service/models"
)

func GetSuggestions(userID string) []models.Product {
	return GetMostSearchedProducts()
}
