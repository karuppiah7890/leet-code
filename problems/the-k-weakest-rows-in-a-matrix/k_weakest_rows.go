package k_weakest_rows

// Cases to consider

// Questions to consider

// Thoughts
// 1. How to optimize? Hmm
// 2. We can use some constraints as a good optimizing condition. Like,
//    - 1 being in front of all 0s
// 	  - matrix[i][j] is either 0 or 1

func KWeakestRows(mat [][]int, k int) []int {
	return KWeakestRowsBruteForce(mat, k)
}

type Row struct {
	index        int
	soldierCount int
}

type KWeakest struct {
	kWeakest []Row
}

func (kWeakest *KWeakest) AddInSortedOrder(newRow Row) {
	if len(kWeakest.kWeakest) == 0 {
		kWeakest.kWeakest = append(kWeakest.kWeakest, newRow)
		return
	}

	insertionPoint := len(kWeakest.kWeakest)

	for index, row := range kWeakest.kWeakest {
		if newRow.soldierCount < row.soldierCount {
			insertionPoint = index
			break
		}

		if newRow.soldierCount == row.soldierCount && newRow.index < row.index {
			insertionPoint = index
			break
		}
	}

	if insertionPoint == len(kWeakest.kWeakest) {
		kWeakest.kWeakest = append(kWeakest.kWeakest, newRow)
		return
	}

	// insert dummy row at the end
	dummyRow := Row{}
	kWeakest.kWeakest = append(kWeakest.kWeakest, dummyRow)

	for index := len(kWeakest.kWeakest) - 1; index > insertionPoint; index-- {
		kWeakest.kWeakest[index] = kWeakest.kWeakest[index-1]
	}

	kWeakest.kWeakest[insertionPoint] = newRow
}

func (kWeakest *KWeakest) Take(k int) []int {
	result := make([]int, k)
	for count := 0; count < k; count++ {
		result[count] = kWeakest.kWeakest[count].index
	}
	return result
}

func NewKWeakest(length int) *KWeakest {
	return &KWeakest{
		kWeakest: make([]Row, 0, length),
	}
}

func KWeakestRowsBruteForce(mat [][]int, k int) []int {
	kWeakest := NewKWeakest(len(mat))
	for rowIndex := 0; rowIndex < len(mat); rowIndex++ {
		soldierCount := 0
		for columnIndex := 0; columnIndex < len(mat[rowIndex]); columnIndex++ {
			if mat[rowIndex][columnIndex] == 0 {
				break
			}
			soldierCount++
		}

		kWeakest.AddInSortedOrder(Row{index: rowIndex, soldierCount: soldierCount})
	}
	return kWeakest.Take(k)
}

func KWeakestRowsQuick(mat [][]int, k int) []int {
	return nil
}
