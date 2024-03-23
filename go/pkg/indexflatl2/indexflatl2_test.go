package indexflatl2

import (
	"reflect"
	"testing"
)

func TestAddVectors(t *testing.T) {
	index := IndexFlatL2{Dimension: 2}
	vectorsToAdd := [][]float64{{1.0, 2.0}, {3.0, 4.0}}
	index.AddVectors(vectorsToAdd)

	if len(index.vectors) != 2 {
		t.Errorf("Expected 2 vectors, got %d", len(index.vectors))
	}

	if !reflect.DeepEqual(index.vectors, vectorsToAdd) {
		t.Errorf("Vectors were not added correctly")
	}
}

func TestEuclideanDistance(t *testing.T) {
	vectorA := []float64{1, 2}
	vectorB := []float64{4, 6}
	distance, err := EuclideanDistance(vectorA, vectorB)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedDistance := 5.0 // sqrt((3^2) + (4^2))
	if distance != expectedDistance {
		t.Errorf("Expected distance %v, got %v", expectedDistance, distance)
	}
}

func TestSearch(t *testing.T) {
	index := IndexFlatL2{Dimension: 2}
	index.AddVectors([][]float64{{1.0, 2.0}, {3.0, 4.0}, {5.0, 6.0}})
	distances, indexes, err := index.Search([]float64{1.0, 2.0}, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedDistances := []float64{0, 2.8284271247461903} // Distance from (1,2) to itself is 0
	expectedIndexes := []int{0, 1}                        // The closest vectors are the first and second ones added

	if !reflect.DeepEqual(distances, expectedDistances) {
		t.Errorf("Expected distances %v, got %v", expectedDistances, distances)
	}

	if !reflect.DeepEqual(indexes, expectedIndexes) {
		t.Errorf("Expected indexes %v, got %v", expectedIndexes, indexes)
	}

	// Test for incorrect dimension
	_, _, err = index.Search([]float64{1.0}, 2)
	if err == nil {
		t.Errorf("Expected dimension error, got nil")
	}

	// Test for k out of bounds
	_, _, err = index.Search([]float64{1.0, 2.0}, 5)
	if err == nil {
		t.Errorf("Expected 'k is out of bounds' error, got nil")
	}
}
