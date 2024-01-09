golang blurb i am kind of noob


// AddItem add item to customer bill
func AddItem(bill map[string]int, unitConversion map[string]int, item string, unit string) bool {
	decUnit, ok := unitConversion[unit]
	if !ok {
		return false
	}
	bill[item] += decUnit
	return true
}

				...	
solution one 	/ \
				 |


				 // AddItem add item to customer bill
func AddItem(bill map[string]int, units map[string]int, item string, unit string) bool {
	value, exists := units[unit]
	if exists {
		bill[item] += value
		return true
	}
	return false
}


...	
solution two 	/ \
				 |