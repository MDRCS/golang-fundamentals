package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)

	// Arrays
	var first_list [5]int
	var second_list = [5]int{1, 2, 3, 4}
	var other_list = [...]int{1, 2, 3, 4}

	// two dimensions arrays
	var matrix [5][5]int
	// Slices
	// difference between Array and slice, is in slice we don' t mention size
	var xx = []int{10, 20, 30}

	/*
		TIP: [...] makes an array. Using [] makes a slice.
	*/

	var y = []int{1, 5: 4, 6, 2: 30, 3} // you can specify the position of every value

	fmt.Println(first_list, len(second_list), other_list, matrix, xx, y)
	fmt.Println("Hello, world!")

	y = append(y, 4, 5, 6)
	fmt.Println(y, "here is the slice")

	// for slices there is capacity
	fmt.Println(cap(y), "here is the slice")

	/*
		to create an empty slice that already has a length or capacity specified.
		That’s the job of the built-in make function.
		It allows us to specify the type, length, and, optionally, the capacity.

	*/

	z := make([]int, 5) // all 5 cases are initialized with zero value // if we append a new value it will be assigned to the latest case
	z = append(z, 10)   // [0 0 0 0 0 0 10]
	z = append(z, 22)   // [0 0 0 0 0 10 22]
	fmt.Println(z[:4])

	zz := make([]int, 5, 8)   // slice with a length of 5 and capacity of 8
	fmt.Println(cap(zz) == 8) // true

	// array is a static tableau, slice is like dynamic arrays, make is like a queue

	/*
		If you have some starting values, or if a slice’s values aren’t going to change,
		then a slice literal is a good choice:

		Example 3-3. Declaring a slice with default values
		data := []int{2, 4, 6, 8} // numbers we appreciate
		If you have a good idea of how large your slice needs to be, but don’t know what those values will be when you are writing the program, use make.”

	*/

	//  + Slices Share Storage Sometimes
	// When you take a slice from a slice, you are not making a copy of the data.
	// Instead, you now have two variables that are sharing memory.

	var a = []int{1, 2, 3, 4}
	b := a[:3]

	fmt.Println(a, b) // a and b are sharing the same memory cases

	// Example below shows how confusing to share space between slices :
	c := make([]int, 0, 5)    // declaring a slice with 5 in capacity
	c = append(c, 1, 2, 3, 4) // appending 4 values
	d := c[:2]                // [1,2] with capacity of 5 // you should see it as [1, 2, . , ., .]
	e := c[2:]                // [3,4] with 3 capacity // cuz storing values in an array is linear.
	fmt.Println(c, d, e)
	fmt.Println(cap(c), cap(d), cap(e))
	d = append(d, 30, 40, 50) // [1 2 30 40] [1 2 30 40 50] [30 40]
	fmt.Println(c, d, e)
	c = append(c, 60) // [1 2 30 40 60] [1 2 30 40 60] [30 40]
	fmt.Println(c, d, e)
	e = append(e, 70)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)

	// Copy Expr

	s := []int{1, 2, 3, 4}
	j := make([]int, 4)
	num := copy(j, s[:2])
	fmt.Println(j, num)
}
