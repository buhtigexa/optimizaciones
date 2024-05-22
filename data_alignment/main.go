package main

import (
	"fmt"
	"unsafe"
)

type Unaligned struct {
	b1 byte
	i  int64
	b2 byte
}

type Aligned struct {
	b1 byte
	b2 byte
	i  int64
}

func main() {
	fmt.Printf("Size of Unaligned: %d bytes\n", unsafe.Sizeof(Unaligned{}))
	fmt.Printf("Offset of b1 in Unaligned: %d bytes\n", unsafe.Offsetof(Unaligned{}.b1))
	fmt.Printf("Offset of i in Unaligned: %d bytes\n", unsafe.Offsetof(Unaligned{}.i))
	fmt.Printf("Offset of b2 in Unaligned: %d bytes\n", unsafe.Offsetof(Unaligned{}.b2))

	fmt.Printf("Size of Aligned: %d bytes\n", unsafe.Sizeof(Aligned{}))
	fmt.Printf("Offset of b1 in Unaligned: %d bytes\n", unsafe.Offsetof(Aligned{}.b1))
	fmt.Printf("Offset of b2 in Unaligned: %d bytes\n", unsafe.Offsetof(Aligned{}.b2))
	fmt.Printf("Offset of i in Unaligned: %d bytes\n", unsafe.Offsetof(Aligned{}.i))

}
