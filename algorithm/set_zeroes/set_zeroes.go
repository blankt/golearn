package set_zeroes

func setZeroes(matrix [][]int) {

}

func resolveOne(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	var (
		line = make([]int, 0, len(matrix)*len(matrix[0]))
		col  = make([]int, 0, len(matrix)*len(matrix[0]))
	)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				line = append(line, i)
				break
			}
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		for i := 0; i < len(matrix); i++ {
			if matrix[i][j] == 0 {
				col = append(col, j)
				break
			}
		}
	}

	for _, v := range line {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[v][j] = 0
		}
	}
	for _, v := range col {
		for j := 0; j < len(matrix); j++ {
			matrix[j][v] = 0
		}
	}
	return
}
