package sudoku

func Solve(data [][]int) [][]int {
	solver := newSudoku(data)
	return solver.solve()
}

func deepcopy(src [][]int) [][]int {
	dest := make([][]int, len(src))
	for i, v := range src {
		dest[i] = make([]int, len(v))
		copy(dest[i], v)
	}
	return dest
}

func count(nums []int, n int) int {
	ret := 0
	for _, v := range nums {
		if v == n {
			ret++
		}
	}
	return ret
}

func check(nums []int) bool {
	for i := 1; i<10; i++ {
		if (count(nums, i) != 1) {
			return false
		}
	}
	return true
}

func duplicate(nums []int) bool {
	for i := 1; i<10; i++ {
		if (count(nums, i) > 1) {
			return true
		}
	}
	return false
}

type Sudoku struct {
	data [][]int
	cells []*Cell
}

func newSudoku(data [][]int) *Sudoku {
	sudoku := new(Sudoku)
	sudoku.data = data
	return sudoku
}

func(self *Sudoku) extractRow(n int) []int {
	return self.data[n - 1]
}

func(self *Sudoku) extractCol(n int) []int {
	ret := make([]int, 9)
	for i, v := range self.data {
		ret[i] = v[n - 1]
	}
	return ret
}

func(self *Sudoku) extractGrid(n int) []int {
	fromRow := (n - 1) / 3 * 3
	fromCol := (n - 1) % 3 * 3

	return self.extractGridByPos(fromCol + 1, fromRow + 1)
}

func(self *Sudoku) extractGridByPos(x, y int) []int {
  fromRow := (y - 1) / 3 * 3
  fromCol := (x - 1) / 3 * 3

  data := deepcopy(self.data)
  a1 := data[fromRow][fromCol:fromCol + 3]
  a2 := data[fromRow + 1][fromCol:fromCol + 3]
  a3 := data[fromRow + 2][fromCol:fromCol + 3]

  return append(a1, append(a2, a3...)...)
}

func(self *Sudoku) isSolved() bool {
	for i := 1; i<10; i++ {
		if !check(self.extractRow(i)) {return false}
		if !check(self.extractCol(i)) {return false}
		if !check(self.extractGrid(i)) {return false}
	}
	return true
}

func(self *Sudoku) isValid() bool {
	for y:=1; y<10; y++ {
		for x:=1;x<10; x++ {
			if !self.calcCell(x, y).isValid() {
				return false
			}
		}
	}
	for i := 1; i<10; i++ {
		if duplicate(self.extractRow(i)) {return false}
		if duplicate(self.extractCol(i)) {return false}
		if duplicate(self.extractGrid(i)) {return false}
	}
	return true
}

func(self *Sudoku) isInvalid() bool {
	return !self.isValid()
}

func(self *Sudoku) getValue(x, y int) int {
	return self.data[y - 1][x - 1]
}

func(self *Sudoku) isAllowed(x, y, v int) bool {
	return count(self.extractRow(y), v) == 0 &&
	       count(self.extractCol(x), v) == 0 &&
	       count(self.extractGridByPos(x, y), v) == 0
}

func(self *Sudoku) calcCell(x, y int) *Cell {
	v := self.getValue(x, y)
	if (v != 0) {
		return newCell(x, y, []int {v})
	} else {
		values := make([]int, 0)
		for i:=1; i<10; i++ {
			if (self.isAllowed(x, y, i)) {
				values = append(values, i)
			}
		}
		return newCell(x, y, values)
	}
}

func(self *Sudoku) clone(x, y, v int) *Sudoku {
	newData := deepcopy(self.data)
	newData[y - 1][x - 1] = v
	return newSudoku(newData)
}

func(self *Sudoku) calcNext() []*Sudoku {
	if self.isSolved() {
		return []*Sudoku{self}
	} else if self.isInvalid() {
		return []*Sudoku{}
	}
	var backtrackCell *Cell = nil
	changed := false
	newData := deepcopy(self.data)
	for y:=1; y<10; y++ {
		for x:=1; x<10; x++ {
			cell := self.calcCell(x, y)
			if (cell.toInt() != newData[y-1][x-1]) {
				newData[y-1][x-1] = cell.toInt()
				changed = true
			}
			if cell.countValues() == 2 && backtrackCell == nil {
				backtrackCell = cell
			}
		}
	}
	if changed {
		return newSudoku(newData).calcNext()
	}
	ret := make([]*Sudoku, 0)
	for _, v := range backtrackCell.values {
		ret = append(ret, self.clone(backtrackCell.x, backtrackCell.y, v).calcNext()...)
	}
	return ret
}

func(self *Sudoku) solve() [][]int {
	list := self.calcNext()
	if len(list) == 1 {
		return list[0].data
	} else {
		return self.data
	}
}

type Cell struct {
	x int
	y int
	values []int
}

func(self *Cell) toInt() int {
	if len(self.values) == 1 {
		return self.values[0]
	} else {
		return 0
	}
}

func(self *Cell) isValid() bool {
	return len(self.values) > 0 
}

func(self *Cell) countValues() int {
	return len(self.values)
}

func newCell(x int, y int, values []int) *Cell {
	ret := new(Cell)
	ret.x = x
	ret.y = y
	ret.values = values
	return ret
}

