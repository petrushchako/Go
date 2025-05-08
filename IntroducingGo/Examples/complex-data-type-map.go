package main

import "fmt"

func main() {
	// 1. Creating a map
	//    - Maps are created using the `map` keyword, followed by the key type in square brackets,
	//      and then the value type.
	//    - Maps are similar to dictionaries in Python or objects in JavaScript.
	//    - The zero value of a map is nil.  A nil map has no keys and cannot be assigned to.
	//    - Maps are typically used to store key-value pairs where you want to look up values
	//      efficiently by key.

	// Example 1: Creating an empty map of string keys and int values.
	var ages map[string]int
	fmt.Println("1. Empty map (nil):", ages) // Output: map[]

	// Example 2: Creating and initializing a map using a map literal.
	//    - A map literal is similar to a struct literal, but the keys are specified
	//      before the values.
	//    - The comma after the last key-value pair is optional.
	ages2 := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Charlie": 35, // The trailing comma is allowed.
	}
	fmt.Println("2. Initialized map:", ages2) // Output: map[Alice:30 Bob:25 Charlie:35]

	// Example 3: Creating a map using the make() function.
	//    - The make() function is the preferred way to create maps.
	//    - You can optionally specify the initial capacity of the map as the second argument.
	//      This can improve performance if you know the approximate number of elements
	//      you'll be storing in the map.
	ages3 := make(map[string]int)
	fmt.Println("3. Map created with make():", ages3) // Output: map[]

	// 2. Adding and accessing map elements
	//    - You can add or update key-value pairs using the assignment operator.
	//    - You can access the value associated with a key using the key in square brackets.

	ages3["Alice"] = 30 // Add "Alice": 30
	ages3["Bob"] = 25   // Add "Bob": 25
	fmt.Println("4. Map after adding elements:", ages3) // Output: map[Alice:30 Bob:25]

	aliceAge := ages3["Alice"] // Access the value associated with the key "Alice".
	fmt.Println("5. Alice's age:", aliceAge)             // Output: 30

	// 3. Checking if a key exists
	//    - When you try to access a key that doesn't exist in the map, Go returns the
	//      zero value of the value type (0 for int, "" for string, false for bool,
	//      nil for pointers, interfaces, and maps).
	//    - To check if a key exists, you can use the following idiom:
	//      value, ok := map[key]
	//    - If the key exists, value is the associated value, and ok is true.
	//    - If the key doesn't exist, value is the zero value, and ok is false.

	ageOfDavid := ages3["David"]           // Key "David" does not exist.
	fmt.Println("6. David's age:", ageOfDavid) // Output: 0 (zero value for int)

	ageOfDavid2, ok := ages3["David"] // Check if "David" exists.
	if ok {
		fmt.Println("7. David's age:", ageOfDavid2)
	} else {
		fmt.Println("7. David's age not found") // Output: David's age not found
	}

	ageOfAlice2, ok := ages3["Alice"] // Check if "Alice" exists.
	if ok {
		fmt.Println("8. Alice's age:", ageOfAlice2) // Output: Alice's age: 30
	} else {
		fmt.Println("8. Alice's age not found")
	}

	// 4. Deleting map elements
	//    - You can delete a key-value pair from a map using the delete() function.
	//    - The delete() function takes the map and the key as arguments.
	//    - It doesn't return any value.
	//    - It's safe to call delete() with a key that doesn't exist; it simply does nothing.

	delete(ages3, "Bob")             // Delete "Bob": 25
	fmt.Println("9. Map after deleting Bob:", ages3) // Output: map[Alice:30]

	delete(ages3, "Bob")             // Delete "Bob" again (no error).
	fmt.Println("10. Map after deleting Bob again:", ages3)

	// 5. Iterating over maps using a for...range loop
	//    - You can iterate over the key-value pairs in a map using a for...range loop.
	//    - The loop variables are the key and the value.
	// 	  - The order of iteration is not specified and is not guaranteed to be the same from one iteration to the next.
	fmt.Println("11. Iterating over map:")
	for name, age := range ages3 {
		fmt.Printf("  %s is %d years old\n", name, age)
	}
	//Expected output:
	//  Alice is 30 years old

	// 6.  Maps are references to a hash table
	//    - When you assign a map to a new variable, both variables refer to the same
	//      underlying hash table.
	//    - If you modify the map through one variable, the changes are visible through
	//      the other variable.
	// 	  - If you pass a map to a function, the function receives a copy of the reference,
	// 	    so any changes the function makes to the map will be visible to the caller.

	ages4 := ages3       // ages4 and ages3 now refer to the same map.
	ages4["Charlie"] = 35 // Add "Charlie": 35 to the map.
	fmt.Println("12. ages3:", ages3) // Output: map[Alice:30 Charlie:35]
	fmt.Println("13. ages4:", ages4) // Output: map[Alice:30 Charlie:35]

	// 7.  Comparing maps
	//     - Maps can only be compared to nil.
	//     - To compare two maps for equality, you must iterate over the key-value pairs
	//       in one map and check if the other map contains the same key-value pairs.

	var map1 map[string]int
	var map2 map[string]int

	fmt.Println("14. map1 == nil:", map1 == nil) // Output: true
	fmt.Println("15. map2 == nil:", map2 == nil) // Output: true

	//Example of comparing two maps:
	map5 := map[string]int{"a": 1, "b": 2}
	map6 := map[string]int{"a": 1, "b": 2}
	map7 := map[string]int{"a": 1, "b": 3}

	fmt.Println("16. Are map5 and map6 equal?", mapsEqual(map5, map6)) //true
	fmt.Println("17. Are map5 and map7 equal?", mapsEqual(map5, map7)) //false
}
// mapsEqual compares two maps for equality.
func mapsEqual(map1, map2 map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key1, value1 := range map1 {
		value2, ok := map2[key1]
		if !ok || value1 != value2 {
			return false
		}
	}
	return true
}
