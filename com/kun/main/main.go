package main

import (
	"alg_go/com/kun/alg/sort"
	"fmt"
	"math/rand"
)

var size int
var array []int

func init() {
	size = 1000

	array = make([]int, size, size)
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(size)
	}
}



func main() {
	fmt.Println(array)
	//sort.BubbleSort(array)
	//sort.SelectSort(array)
	//sort.InsertSort(array)
	sort.MergeSort(array)
	fmt.Println(array)

}
