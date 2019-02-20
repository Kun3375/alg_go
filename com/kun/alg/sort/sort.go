package sort

import (
	"math/rand"
)

func BubbleSort(slice []int) {

	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

func SelectSort(slice []int) {
	if len(slice) == 0 {
		return
	}
	var temp int
	var tempIndex int
	for i := 0; i < len(slice); i++ {
		tempIndex = i
		temp = slice[i]
		for j := i; j < len(slice); j++ {
			if slice[j] > temp {
				tempIndex = j
				temp = slice[j]
			}
		}
		slice[i], slice[tempIndex] = temp, slice[i]
	}
}

func InsertSort(slice []int) {
	if len(slice) == 0 {
		return
	}
	var tempIndex int
	var tempHolder int
	for i := 0; i < len(slice); i++ {
		tempIndex = i
		tempHolder = slice[i]
		for ; tempIndex > 0 && slice[tempIndex-1] > tempHolder; tempIndex-- {
			slice[tempIndex] = slice[tempIndex-1]
		}
		slice[tempIndex] = tempHolder
	}
}

func MergeSort(slice []int) {
	mergeSortCore(slice)
}

func mergeSortCore(slice []int) {
	if len(slice) < 16 {
		InsertSort(slice)
		return
	}
	mid := len(slice) >> 1
	mergeSortCore(slice[:mid])
	mergeSortCore(slice[mid:])
	mergeTwoPart(slice)
}

func mergeTwoPart(slice []int) {
	start, end := 0, len(slice)
	mid := end >> 1

	if slice[mid-1] < slice[mid] {
		return
	}
	if slice[end-1] < slice[start] {
		temp := make([]int, mid-start, mid-start)
		copy(temp, slice[mid:end])
		copy(slice[mid:end], slice[start:mid])
		copy(slice[start:mid], temp)
		return
	}
	temp := make([]int, end-start, end-start)
	tempIndex, i, j := 0, start, mid
	for i < mid && j < end {
		if slice[i] < slice[j] {
			temp[tempIndex] = slice[i]
			tempIndex++
			i++
		} else {
			temp[tempIndex] = slice[j]
			tempIndex++
			j++
		}
	}
	for i < mid {
		temp[tempIndex] = slice[i]
		tempIndex++
		i++
	}
	for j < end {
		temp[tempIndex] = slice[j]
		tempIndex++
		j++
	}
	copy(slice[start:end], temp)
}

func QuickSort(slice []int) {
	if len(slice) < 2 {
		return
	}
	midIndex := quickSortPartition(slice)
	QuickSort(slice[:midIndex])
	QuickSort(slice[midIndex+1:])
}

func quickSortPartition(slice []int) int {
	rIndex := rand.Intn(len(slice))
	slice[0], slice[rIndex] = slice[rIndex], slice[0]

	temp := slice[0]
	i, j := 0, 1
	for j < len(slice) {
		if slice[j] < temp {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
		j++
	}

	slice[0], slice[i] = slice[i], slice[0]
	return i
}

func QuickSort2Way(slice []int) {
	if len(slice) <= 16 {
		InsertSort(slice)
		return
	}
	midIndex := quickSort2WayPartition(slice)
	QuickSort2Way(slice[:midIndex])
	QuickSort2Way(slice[midIndex+1:])
}

func quickSort2WayPartition(slice []int) int {
	rIndex := rand.Intn(len(slice))
	slice[0], slice[rIndex] = slice[rIndex], slice[0]

	temp := slice[0]
	i, j := 1, len(slice)-1
	for true {
		for i < j && slice[i] < temp {
			i++
		}
		for j > i && slice[j] > temp {
			j--
		}
		if i >= j {
			break
		}
		slice[i], slice[j] = slice[j], slice[i]
		i++
		j--
	}
	slice[0], slice[j] = slice[j], slice[0]
	return j
}

func QuickSort3Way(slice []int) {
	if len(slice) < 16 {
		InsertSort(slice)
		return
	}

	eq, lt := quickSort3WayPartition(slice)
	QuickSort3Way(slice[lt:])
	QuickSort3Way(slice[:eq])
}

func quickSort3WayPartition(slice []int) (int, int) {
	rIndex := rand.Intn(len(slice))
	slice[0], slice[rIndex] = slice[rIndex], slice[0]

	i, e, j := 1, 1, len(slice)
	for i < j {
		if slice[0] > slice[i] {
			slice[e], slice[i] = slice[i], slice[e]
			e++
			i++
			continue
		}
		if slice[0] < slice[i] {
			j--
			slice[j], slice[i] = slice[i], slice[j]
			continue
		}
		i++
	}
	e--
	slice[0], slice[e] = slice[e], slice[0]
	return e, j
}
