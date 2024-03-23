package indexflatl2

import (
	"errors"
	"math"
	"sort"
)

// IndexFlatL2 represents a structure for storing vectors and performing searches on them.
type IndexFlatL2 struct {
	Dimension int
	vectors   [][]float64
}

// IndexValuePair holds a pair of a distance and an index in the vectors slice.
type IndexValuePair struct {
	Distance float64
	Index    int
}

// AddVectors appends new vectors to the IndexFlatL2 instance.
func (index *IndexFlatL2) AddVectors(vectors [][]float64) {
	index.vectors = append(index.vectors, vectors...)
}

// Search performs a search using Euclidean distance, returning the distances and indexes of the k nearest vectors.
func (index *IndexFlatL2) Search(inferenceVector []float64, k int) ([]float64, []int, error) {
	if len(inferenceVector) != index.Dimension {
		return nil, nil, errors.New("incorrect dimensions")
	}

	if k > len(index.vectors) {
		return nil, nil, errors.New("k is out of bounds")
	}

	distances := make([]IndexValuePair, len(index.vectors))
	for i, vector := range index.vectors {
		distance, err := EuclideanDistance(inferenceVector, vector)
		if err != nil {
			return nil, nil, err
		}
		distances[i] = IndexValuePair{distance, i}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	kNNDistances := make([]float64, k)
	kNNElementIndex := make([]int, k)
	for i := 0; i < k; i++ {
		kNNDistances[i] = distances[i].Distance
		kNNElementIndex[i] = distances[i].Index
	}

	return kNNDistances, kNNElementIndex, nil
}

// EuclideanDistance calculates the Euclidean distance between two vectors.
func EuclideanDistance(a, b []float64) (float64, error) {
	if len(a) != len(b) {
		return 0, errors.New("vectors must have the same dimension")
	}
	var distanceSquared float64
	for i := range a {
		distanceSquared += math.Pow(a[i]-b[i], 2)
	}
	return math.Sqrt(distanceSquared), nil
}
