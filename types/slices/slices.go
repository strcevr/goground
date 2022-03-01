package slices

import (
	"fmt"
	"unsafe"
)

func ReallocatedAndAssignCases () {

	a := []int{1, 2, 3, 4, 5}
	func(b []int) {
		b[1] = 7
		b = append(b, 1)
		fmt.Println(b) // prints [1 7 3 4 5 1].
	}(a)
	fmt.Println(a) // prints [1 7 3 4 5].

	c := a[0:4]
	fmt.Println(c) // prints [1 7 3 4]
	a = a[:0] // Slice the slice to zero length to keep the underlying array and allocated memory.
	fmt.Println(a) // Prints []

	fmt.Println("Try to assign value into already empty slice")
	func(){
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Recovered: ", err)
			}
		}()
		a[2] = 3 // Panic
	}()
	fmt.Println(a) // Prints []
	fmt.Println(a[:2]) // If the slice is extended again, the original data reappears. Prints [1 7]

	var a2 []int
	a2 = append(a2,20)
	fmt.Println(a2) // Prints [20]
}

func SizesCases () {
	// 3 words by default.
	var emptySlice []string
	fmt.Println(unsafe.Sizeof(emptySlice)) // prints 24.
	// 3 words by default.
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(a)) // prints 24.
	// Array, size = length * element_size.
	var b [5]bool
	fmt.Println(unsafe.Sizeof(b)) // prints 5.
	//
	c := [...]int{1, 2, 3, 4}
	fmt.Println(unsafe.Sizeof(c)) // prints 32.
}

func CopyCases () {
	var dst []int
	src := []int{3, 2, 1}
	copy(dst, src)
	fmt.Println(dst)
	dst = append(dst, 1)
	fmt.Println(dst)
	copy(dst, src)
	fmt.Println(dst)
	copy(dst, []int{})
	fmt.Println(dst)
}

func ReturnNilSlice () {
	a := func() []*int {
			return nil
		}()
	fmt.Println(len(a))
	fmt.Println(a)
}