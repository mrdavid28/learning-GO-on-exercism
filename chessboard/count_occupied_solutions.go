//count occupied solutions  ##

// ##1 first solution
func CountOccupied(cb Chessboard) int {
	var occupiedSquares int = 0
	for file := range cb {
		for _, isOccupied := range cb[file] {
			if isOccupied {
				occupiedSquares++
			}
		}
	}
	return occupiedSquares
}

// ##2 second solution
func CountOccupied(cb Chessboard) (count int) {
	for _, rank := range cb {
		for _, occupied := range rank {
			if occupied {
				count++
			}
		}
	}
	return count
}

// ##third solution
func CountOccupied(cb Chessboard) int {
	count := 0
	for k, _ := range cb {
		count += CountInRank(cb, k)
	}
	return count
}

// fourth solution
func CountOccupied(cb Chessboard) int {
	occupied := 0
	for i, _ := range cb {
		for _, sqOccupied := range cb[i] {
			if sqOccupied {
				occupied++
			}
		}
	}
	return occupied
}