package main

func main() {

}

func createMatrix(size int) [][]int64 {
	matrix := make([][]int64, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int64, size)
	}
	return matrix
}
