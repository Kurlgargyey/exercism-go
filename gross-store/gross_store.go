package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{"quarter_of_a_dozen": 3, "half_of_a_dozen": 6, "dozen": 12, "small_gross": 120, "gross": 144, "great_gross": 1728}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unit_value, exists := units[unit]
	if exists {
		bill[item] += unit_value
		return true
	} else {
		return false
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, item_exists := bill[item]
	quan, unit_exists := units[unit]
	if item_exists && unit_exists {
		switch {
		case bill[item] > quan:
			bill[item] -= quan
		case bill[item] == quan:
			delete(bill, item)
		default:
			return false
		}
		return true

	} else {
		return false
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quan, exists := bill[item]
	return quan, exists
}
