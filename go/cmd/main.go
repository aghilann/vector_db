package main

import (
	"math/rand"
)

// GenerateRandomArray generates a random array of size n with values between 0 and 1
func GenerateRandomArray(n int) []float64 {
	array := make([]float64, n)
	for i := 0; i < n; i++ {
		array[i] = rand.Float64() // Generates a random float between 0 and 1
	}
	return array
}

// GenerateRandom2DArray generates a 2D list of size rows x cols with random values between 0 and 1
func GenerateRandom2DArray(rows, cols int) [][]float64 {
	arraySize := rows * cols
	randomArray := GenerateRandomArray(arraySize)
	twoDArray := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		twoDArray[i] = randomArray[i*cols : (i+1)*cols] // Slice the 1D array to create rows
	}
	return twoDArray
}

func main() {
	const dimension = 10
	index := IndexFlatL2{Dimension: dimension}
	random2DArray := GenerateRandom2DArray(10, dimension)
	sampleInference := GenerateRandom2DArray(1, dimension)[0]

	index.AddVectors(random2DArray)
	index.Search(sampleInference, 2)
}
