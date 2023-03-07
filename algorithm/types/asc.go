package types

import "sort"

type Array []int32

type ArrayRequest struct {
	Array []int `json:"array" binding:"required"`
	Int   int   `json:"integer" binding:"required"`
}

type IntRequestResponse struct {
	Int int `json:"integer" binding:"required"`
}

func (arr *ArrayRequest) SortInts() *IntRequestResponse {
	sort.Ints(arr.Array)
	return &IntRequestResponse{
		Int: SliceIndex(len(arr.Array), func(i int) bool { return arr.Array[i] == arr.Int }),
	}
}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
