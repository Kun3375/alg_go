package sort

import (
	"math/rand"
	"testing"
)

func BenchmarkMergeSort(b *testing.B) {
	b.N = 100
	prepare(MergeSort)
}

func BenchmarkQuickSort(b *testing.B) {
	b.N = 100
	prepare(QuickSort)
}

func BenchmarkQuickSort2Way(b *testing.B) {
	b.N = 100
	prepare(QuickSort2Way)
}

func BenchmarkQuickSort3Way(b *testing.B) {
	b.N = 100
	prepare(QuickSort3Way)
}

func BenchmarkSelectSort(b *testing.B) {
	b.N = 100
	prepare(SelectSort)
}

func BenchmarkInsertSort(b *testing.B) {
	b.N = 100
	prepare(InsertSort)
}

func BenchmarkBubbleSort(b *testing.B) {
	b.N = 100
	prepare(BubbleSort)
}

func prepare(f func([]int)) {
	size := 100000
	ints := make([]int, size, size)
	for i := 0; i < len(ints); i++ {
		ints[i] = rand.Intn(size)
	}
	f(ints)
}
