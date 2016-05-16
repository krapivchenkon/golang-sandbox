// Here we test various ways of passing slices to functions
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Helper function reads 2-dimensional matrix into slice of slices
func readMatrix(s *bufio.Scanner, arr [][]int32, N *int32) {
	var row, ind int32 = -1, 0
	for s.Scan() {
		if ind%*N == 0 {
			row++
		}
		ind++
		i, err := strconv.ParseInt(s.Text(), 10, 32)
		if err != nil {
			panic(err)
		}
		arr[row] = append(arr[row], int32(i))
	}

}

func readArray(s *bufio.Scanner, arr []int32, N *int32) {
	// var row, ind int32 = -1, 0
	for i := int32(0); i < *N && s.Scan(); i++ {
		el, err := strconv.ParseInt(s.Text(), 10, 32)
		fmt.Println("el:", el)
		if err != nil {
			panic(err)
		}
		arr[i] = int32(el)
	}
	fmt.Println("changed slice:", arr)

}

// extend slice when its passed as value
func readArraySliceAsValue(s *bufio.Scanner, arr []int32, N *int32) {
	// var row, ind int32 = -1, 0
	for i := int32(0); i < *N && s.Scan(); i++ {
		el, err := strconv.ParseInt(s.Text(), 10, 32)
		fmt.Println("el:", el)
		if err != nil {
			panic(err)
		}
		// arr[i] = int32(el)
		arr = append(arr, int32(el))
	}
	fmt.Println("changed slice:", arr)

}

func readArraySliceNew(s *bufio.Scanner, arr *[]int32, N *int32) {
	// var row, ind int32 = -1, 0
	for i := int32(0); i < *N && s.Scan(); i++ {
		el, err := strconv.ParseInt(s.Text(), 10, 32)
		fmt.Println("el:", el)
		if err != nil {
			panic(err)
		}
		// arr[i] = int32(el)
		*arr = append(*arr, int32(el))
	}
	// arr is a pointer to slice preallocated with new
	fmt.Println("changed slice:", *arr)

}

func readArraySliceAsReference(s *bufio.Scanner, arr *[]int32, N *int32) {
	// var row, ind int32 = -1, 0
	for i := int32(0); i < *N && s.Scan(); i++ {
		el, err := strconv.ParseInt(s.Text(), 10, 32)
		fmt.Println("el:", el)
		if err != nil {
			panic(err)
		}
		// arr[i] = int32(el)
		*arr = append(*arr, int32(el))
	}
	// arr is a pointer to slice preallocated with new
	fmt.Println("changed slice:", *arr)

}

func readArraySliceAsReferenceNoPrealloc(s *bufio.Scanner, arr *[]int32, N *int32) {
	// var row, ind int32 = -1, 0
	for i := int32(0); i < *N && s.Scan(); i++ {
		el, err := strconv.ParseInt(s.Text(), 10, 32)
		fmt.Println("el:", el)
		if err != nil {
			panic(err)
		}
		// arr[i] = int32(el)
		*arr = append(*arr, int32(el))
	}
	// arr is a pointer to slice preallocated with new
	fmt.Println("changed slice:", *arr)

}

// cat following lines to this script:
// 6
// 1 2 3 4 10 11
// 1 2 3 4 10 11
// 1 2 3 4 10 11
// 1 2 3 4 10 11
// 1 2 3 4 10 11
func main() {

	// SAMPLE: read input sample using bufio Scanner
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	scan.Scan()
	n, _ := strconv.ParseInt(scan.Text(), 10, 32)
	var N int32 = int32(n)

	arr := make([]int32, N)
	readArray(scan, arr, &N)
	fmt.Println("original slice:", arr)

	arr2 := make([]int32, N)
	// slice itself is a strusture of pointers that actually point to underlining array element and length
	// so here we pass it as a value(copy)
	// so when we append values - we change copy of initial slice and not the original
	readArraySliceAsValue(scan, arr2, &N)
	// output here will be. NOTE: slice is preallocated with make function
	// changed slice: [0 0 0 0 0 0 1 2 3 4 10 11]
	// original slice: [0 0 0 0 0 0]
	//
	// Below is the same case as previous but array created with new
	fmt.Println("original slice:", arr2)

	arr3 := new([]int32)
	// in this case we pass slice as reference to nil slice as its created with new
	// so when we append values - we appendind in to the same original slice
	readArraySliceNew(scan, arr3, &N)
	// output here will be. NOTE: slice is preallocated with make function
	// changed slice: [1 2 3 4 10 11]
	// original slice: [1 2 3 4 10 11]
	//
	//  arr3 is a pointer to slice
	fmt.Println("original slice:", *arr3)

	arr4 := make([]int32, N)
	// in this sample we preallocate slice with N values but pass reference to slice to a function
	readArraySliceAsReference(scan, &arr4, &N)
	// output here will be. NOTE: slice is preallocated with make function
	// changed slice: [0 0 0 0 0 0 1 2 3 4 10 11]
	// original slice: [0 0 0 0 0 0 1 2 3 4 10 11]
	//
	fmt.Println("original slice:", arr4)

	var arr5 []int32
	fmt.Println(arr5)
	// in this sample we preallocate slice with N values but pass reference to slice to a function
	readArraySliceAsReferenceNoPrealloc(scan, &arr5, &N)
	// output here will be. NOTE: slice is preallocated with make function
	// changed slice: [1 2 3 4 10 11]
	// original slice: [1 2 3 4 10 11]
	//
	fmt.Println("original slice:", arr5)

	// When we don't return changed slice in a function - it should be passed to function by reference
	// in case we return changed slice from a function - we can pass it as a copy value
}
