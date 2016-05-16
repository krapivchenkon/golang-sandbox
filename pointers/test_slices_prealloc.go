// Here we test various ways of passing slices to functions
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func readArray(src []int32, arr []int32, N *int32) {
	// var row, ind int32 = -1, 0
	for i := int32(0); i < *N; i++ {
		arr[i] = src[i]
	}
	// fmt.Println("changed slice:", arr)

}

func readArraySliceNew(src []int32, arr *[]int32, N *int32) {
	// var row, ind int32 = -1, 0
	for i := int32(0); i < *N; i++ {
		*arr = append(*arr, src[i])
		// *arr = append(*arr, int32(el))
	}

	// arr is a pointer to slice preallocated with new
	// fmt.Println("changed slice:", *arr)

}

func readArraySliceAsReferenceNoPrealloc(src []int32, arr *[]int32, N *int32) {
	// var row, ind int32 = -1, 0
	for i := int32(0); i < *N; i++ {
		*arr = append(*arr, src[i])
		// *arr = append(*arr, int32(el))
	}
	// arr is a pointer to slice preallocated with new
	// fmt.Println("changed slice:", *arr)

}

func main() {

	Tests := []int32{10, 1000, 5000, 10000, 50000, 100000, 500000, 1000000, 5000000}
	for test, N := range Tests {
		fmt.Printf("N=%d\n", N)
		// var N int32 = 10000000
		src := make([]int32, N, N)
		for i := int32(0); i < N; i++ {
			src[i] = rand.Int31()
		}
		// fmt.Println("source:", src)

		// fmt.Println("Test"+i, "--------Prealloc fixed")
		now := time.Now()
		arr := make([]int32, N)
		// Here we preallocate values and then set them to array
		readArray(src, arr, &N)
		// fmt.Println("original slice:", arr)
		fmt.Println("Test#", test, "--------Prealloc fixed      -->time:", time.Now().Sub(now))
		// fmt.Println("--------Nil slice New append")
		now = time.Now()
		arr3 := new([]int32)
		// in this case we pass slice as reference to nil slice as its created with new
		// so when we append values - we appendind in to the same original slice
		readArraySliceNew(src, arr3, &N)
		// fmt.Println("original slice:", *arr3)
		fmt.Println("Test#", test, "--------Nil slice New append-->time:", time.Now().Sub(now))
		// fmt.Println("time:", time.Now().Sub(now))

		// fmt.Println("--------Nil slice append")
		now = time.Now()
		var arr5 []int32
		// in this sample we preallocate slice with N values but pass reference to slice to a function
		readArraySliceAsReferenceNoPrealloc(src, &arr5, &N)
		// fmt.Println("original slice:", arr5)
		fmt.Println("Test#", test, "--------Nil slice append    -->time:", time.Now().Sub(now))
		// fmt.Println("time:", time.Now().Sub(now))
		// When we don't return changed slice in a function - it should be passed to function by reference
		// in case we return changed slice from a function - we can pass it as a copy value
	}
}
