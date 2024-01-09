package main

import (
	"fmt"
)

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	var NewBill map[string]int = make(map[string]int)
	return NewBill
}

// AddItem adds an item to customer bill.
func AddItem(bill map[string]int, units map[string]int, item string, unit string) bool {
	value, exists := units[unit]
	if exists {
		bill[item] += value
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	bill_value, exists := bill[item]

	if !exists {
		return false
	}
	unit_value, is_ok := units[unit]
	if !is_ok {
		return false
	}
	if unit_value > bill_value {
		return false
	}
	if bill_value == unit_value {
		delete(bill, item)
		return true
	}
	bill[item] -= unit_value
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, ok := bill[item]
	if ok == true {
		return qty, true
	}
	return 0, false
}

func main() {
	bill := NewBill()
	units := Units()
	ok := AddItem(bill, units, "carrot", "gross")
	okk := AddItem(bill, units, "cucumber", "dozen")
	fmt.Println(ok)
	fmt.Println(okk)
	fmt.Println(bill["cucumber"])
}
