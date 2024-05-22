package main

type StructOfSlice struct {
	a []int64
	b []int64
}

func sumStructOfSlice(bar StructOfSlice) int64 {
	var total int64
	for i := 0; i < len(bar.a); i++ {
		total += bar.a[i]
	}
	return total
}

type Struct struct {
	a int64
	b int64
}

func sumSliceOfStruct(foos []Struct) int64 {
	var total int64
	for i := 0; i < len(foos); i++ {
		total += foos[i].a
	}
	return total
}
