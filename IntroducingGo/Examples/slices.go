package main

import "fmt"

func main() {
	// 1. Creating and initializing slices
	// From an array (creates a slice referencing a portion of the array)
	arr := [5]int{10, 20, 30, 40, 50}
	slice1 := arr[1:4] // Elements at index 1, 2, and 3 (20, 30, 40)
	fmt.Println("Slice 1 (from array):", slice1)
	fmt.Println("Length of slice 1:", len(slice1))
	fmt.Println("Capacity of slice 1:", cap(slice1)) // Capacity is the underlying array size from the starting index

	// Using the make() function (creates a slice with a specified length and capacity)
	slice2 := make([]string, 3) // Length 3, capacity 3
	slice2[0] = "apple"
	slice2[1] = "banana"
	slice2[2] = "cherry"
	fmt.Println("Slice 2 (using make):", slice2)
	fmt.Println("Length of slice 2:", len(slice2))
	fmt.Println("Capacity of slice 2:", cap(slice2))

	slice3 := make([]int, 2, 5) // Length 2, capacity 5
	slice3[0] = 100
	slice3[1] = 200
	fmt.Println("Slice 3 (make with capacity):", slice3)
	fmt.Println("Length of slice 3:", len(slice3))
	fmt.Println("Capacity of slice 3:", cap(slice3))

	// Directly initializing a slice literal
	slice4 := []float64{3.14, 2.71, 1.618}
	fmt.Println("Slice 4 (literal):", slice4)
	fmt.Println("Length of slice 4:", len(slice4))
	fmt.Println("Capacity of slice 4:", cap(slice4))

	fmt.Println("\n--------------------\n")

	// 2. Accessing and modifying slice elements
	fmt.Println("Element at index 1 of slice 4:", slice4[1])
	slice4[0] = 3.0 // Modifying an element
	fmt.Println("Modified slice 4:", slice4)

	fmt.Println("\n--------------------\n")

	// 3. Appending elements to a slice (can increase the slice's length and potentially its capacity)
	slice5 := []int{1, 2}
	fmt.Println("Initial slice 5:", slice5, "Length:", len(slice5), "Capacity:", cap(slice5))

	slice5 = append(slice5, 3)
	fmt.Println("After appending 3:", slice5, "Length:", len(slice5), "Capacity:", cap(slice5))

	slice5 = append(slice5, 4, 5, 6) // Appending multiple elements
	fmt.Println("After appending 4, 5, 6:", slice5, "Length:", len(slice5), "Capacity:", cap(slice5))

	slice6 := []int{7, 8}
	slice5 = append(slice5, slice6...) // Appending another slice (using the "..." spread operator)
	fmt.Println("After appending slice 6:", slice5, "Length:", len(slice5), "Capacity:", cap(slice5))

	fmt.Println("\n--------------------\n")

	// 4. Slicing a slice (creating a new slice that refers to a portion of the original)
	subSlice1 := slice5[2:5] // Elements at index 2, 3, and 4 of slice5
	fmt.Println("Sub-slice 1 of slice 5:", subSlice1, "Length:", len(subSlice1), "Capacity:", cap(subSlice1))
	// Capacity of subSlice1 is from the starting index (2) to the end of slice5's underlying array

	subSlice2 := slice5[:3] // Elements from the beginning up to (but not including) index 3
	fmt.Println("Sub-slice 2 of slice 5:", subSlice2, "Length:", len(subSlice2), "Capacity:", cap(subSlice2))

	subSlice3 := slice5[4:] // Elements from index 4 to the end
	fmt.Println("Sub-slice 3 of slice 5:", subSlice3, "Length:", len(subSlice3), "Capacity:", cap(subSlice3))

	subSlice4 := slice5[:] // A slice referencing the entire underlying array of slice5
	fmt.Println("Sub-slice 4 of slice 5:", subSlice4, "Length:", len(subSlice4), "Capacity:", cap(subSlice4))

	fmt.Println("\n--------------------\n")

	// 5. Copying slices (creates a new slice with copied elements)
	sourceSlice := []int{100, 200, 300}
	destinationSlice := make([]int, len(sourceSlice)) // Create a destination slice of the same length
	numCopied := copy(destinationSlice, sourceSlice)
	fmt.Println("Source slice:", sourceSlice)
	fmt.Println("Destination slice:", destinationSlice)
	fmt.Println("Number of elements copied:", numCopied)

	// Copying to a smaller slice (only the number of elements that fit will be copied)
	smallerDestination := make([]int, 2)
	numCopiedSmaller := copy(smallerDestination, sourceSlice)
	fmt.Println("Smaller destination:", smallerDestination)
	fmt.Println("Number copied to smaller:", numCopiedSmaller)
}
