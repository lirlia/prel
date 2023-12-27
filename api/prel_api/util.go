package api

import (
	"sort"
)

var smallestPageSize int
var largestPageSize int

func init() {
	var pageSize PageSize
	var pageSizes []int
	for _, v := range pageSize.AllValues() {
		pageSizes = append(pageSizes, int(v))
	}

	sort.Slice(pageSizes, func(i, j int) bool {
		return pageSizes[i] < pageSizes[j]
	})

	smallestPageSize = pageSizes[0]
	largestPageSize = pageSizes[len(pageSizes)-1]
}

func SmallestPageSize() int {
	return smallestPageSize
}

func LargestPageSize() int {
	return largestPageSize
}

func AllPageSize() []int {
	var pageSize PageSize
	var pageSizes []int
	for _, v := range pageSize.AllValues() {
		pageSizes = append(pageSizes, int(v))
	}

	return pageSizes
}
