package piscine

import "strings"

type food struct {
	preptime int
}

// Menu to map food items to their preparation time
var menu = map[string]food{
	"burger":  {preptime: 15},
	"chips":   {preptime: 10},
	"nuggets": {preptime: 12},
}

// FoodDeliveryTime function calculates the preparation time for the given order
func FoodDeliveryTime(order string) int {
	// Convert order to lowercase to handle case-insensitivity
	order = strings.ToLower(order)
	// Check if the order is valid in the menu
	if item, exists := menu[order]; exists {
		return item.preptime
	}
	// Return 404 if the item doesn't exist in the menu
	return 404
}
