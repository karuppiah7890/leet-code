package matrix

func setZeroes(matrix [][]int) {
	rowsToBeSetAsZeroes := make(map[int]struct{}, len(matrix))
	columnsToBeSetAsZeroes := make(map[int]struct{}, len(matrix[0]))

	for rowIndex := range matrix {
		for columnIndex := range matrix[rowIndex] {
			if matrix[rowIndex][columnIndex] == 0 {
				rowsToBeSetAsZeroes[rowIndex] = struct{}{}
				columnsToBeSetAsZeroes[columnIndex] = struct{}{}

				makeRowZeroesPartially(matrix, rowIndex, 0, columnIndex)
				makeColumnZeroesPartially(matrix, columnIndex, 0, rowIndex)

				continue
			}

			_, rowIndexFound := rowsToBeSetAsZeroes[rowIndex]

			if rowIndexFound {
				matrix[rowIndex][columnIndex] = 0
				continue
			}

			_, columnIndexFound := columnsToBeSetAsZeroes[columnIndex]

			if columnIndexFound {
				matrix[rowIndex][columnIndex] = 0
				continue
			}
		}
	}
}

func makeRowZeroesPartially(matrix [][]int, rowIndex int, startColumnIndex int, endColumnIndex int) {

	for columnIndex := startColumnIndex; columnIndex <= endColumnIndex; columnIndex++ {
		matrix[rowIndex][columnIndex] = 0
	}
}

func makeColumnZeroesPartially(matrix [][]int, columnIndex int, startRowIndex int, endRowIndex int) {

	for rowIndex := startRowIndex; rowIndex <= endRowIndex; rowIndex++ {
		matrix[rowIndex][columnIndex] = 0
	}
}
