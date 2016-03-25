package sudoku

func Solve(data [][]int) [][]int {
	return data
}


func CheckRow(data [][]int, x int, y int, val int) bool {
	for ix := 0; ix < 9; ix++ {
		if ix == x { 
			continue 
		}
		if data[ix][y] == val { 
			return false 
		}
	}
	return true
}

func CheckColumn(data [][]int, x int, y int, val int) bool {
	for iy := 0; iy < 9; iy++ {
		if iy == y { 
			continue 
		}
		if data[x][iy] == val {
			return false 
		}
	}
	return true
}

func CheckCell(data [][]int, x int, y int, val int) bool {
	sx := x - (x % 3)
	sy := y - (y % 3)

	for ix := sx; ix < sx + 3; ix++ {
		for iy := sy; iy < sy + 3; iy++ {
			if y == iy && x == ix { 
				continue 
			}
			if data[ix][iy] == val { 
				return false 
			}
		}
	}
	return true
}
