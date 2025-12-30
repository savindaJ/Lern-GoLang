package main

import (
	"fmt"
	"reflect"
	"slices"
	"sort"
)

// ============================================================
// 🎯 GO SLICES - COMPREHENSIVE DEEP LEARNING GUIDE
// ============================================================
// A slice is a dynamically-sized, flexible view into the elements
// of an array. Slices are more common than arrays in Go.
//
// Slice Structure (Internal):
// ┌─────────────────────────────────────────┐
// │  Slice Header (24 bytes on 64-bit)      │
// ├─────────────────────────────────────────┤
// │  Pointer to underlying array           │
// │  Length (number of elements)            │
// │  Capacity (max elements before resize)  │
// └─────────────────────────────────────────┘

func main() {
	fmt.Println("=" + "=================================================")
	fmt.Println("🚀 GO SLICES - COMPLETE DEEP LEARNING TUTORIAL")
	fmt.Println("=" + "=================================================")
	fmt.Println()

	// Run all examples
	BasicSliceDeclaration()
	SliceWithMake()
	SliceLengthAndCapacity()
	SlicingOperations()
	AppendingToSlices()
	CopyingSlices()
	NilVsEmptySlice()
	MultiDimensionalSlices()
	SliceIteration()
	SliceSorting()
	SliceSearching()
	SliceFiltering()
	SliceMapping()
	SliceComparison()
	SliceMemoryInternals()
	SliceGotchas()
	SliceTricksAndPatterns()
	ModernSlicePackage()
}

// ============================================================
// 1️⃣ BASIC SLICE DECLARATION
// ============================================================
func BasicSliceDeclaration() {
	fmt.Println("\n📚 1. BASIC SLICE DECLARATION")
	fmt.Println("─" + "───────────────────────────────────────────────")

	// Method 1: Slice literal (most common)
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice literal:", slice1)

	// Method 2: Empty slice literal
	slice2 := []int{}
	fmt.Println("Empty slice literal:", slice2)

	// Method 3: Declare without initializing (nil slice)
	var slice3 []int
	fmt.Println("Nil slice:", slice3, "| Is nil?", slice3 == nil)

	// Method 4: Slice from array
	array := [5]int{10, 20, 30, 40, 50}
	slice4 := array[1:4] // Elements at index 1, 2, 3
	fmt.Println("Slice from array [1:4]:", slice4)

	// Method 5: Slice of strings
	fruits := []string{"Apple", "Banana", "Cherry", "Date"}
	fmt.Println("String slice:", fruits)

	// Method 6: Slice of different types
	floats := []float64{1.1, 2.2, 3.3}
	bools := []bool{true, false, true}
	fmt.Println("Float slice:", floats)
	fmt.Println("Bool slice:", bools)
}

// ============================================================
// 2️⃣ CREATING SLICES WITH make()
// ============================================================
func SliceWithMake() {
	fmt.Println("\n📚 2. CREATING SLICES WITH make()")
	fmt.Println("─" + "───────────────────────────────────────────────")

	// make([]T, length) - creates slice with length and capacity = length
	slice1 := make([]int, 5)
	fmt.Printf("make([]int, 5): %v | len=%d, cap=%d\n", slice1, len(slice1), cap(slice1))

	// make([]T, length, capacity) - creates slice with specified length and capacity
	slice2 := make([]int, 3, 10)
	fmt.Printf("make([]int, 3, 10): %v | len=%d, cap=%d\n", slice2, len(slice2), cap(slice2))

	// When to use make():
	// - When you know the size upfront (better performance)
	// - When you need specific capacity to avoid reallocations
	// - When initializing with zero values is desired

	// Pre-allocating for performance
	slice3 := make([]int, 0, 100) // Empty slice with capacity 100
	fmt.Printf("Pre-allocated: %v | len=%d, cap=%d\n", slice3, len(slice3), cap(slice3))

	// Common mistake: make creates zero-valued elements
	slice4 := make([]int, 5)
	slice4 = append(slice4, 1, 2, 3) // This adds AFTER the 5 zeros!
	fmt.Println("⚠️  After append to make(5):", slice4)

	// Correct way if you want to append from start
	slice5 := make([]int, 0, 5)
	slice5 = append(slice5, 1, 2, 3)
	fmt.Println("✅ Correct way:", slice5)
}

// ============================================================
// 3️⃣ SLICE LENGTH AND CAPACITY
// ============================================================
func SliceLengthAndCapacity() {
	fmt.Println("\n📚 3. SLICE LENGTH AND CAPACITY")
	fmt.Println("─" + "───────────────────────────────────────────────")

	// Length: number of elements in slice
	// Capacity: number of elements in underlying array (from slice start)

	slice := make([]int, 3, 5)
	fmt.Printf("Initial: %v | len=%d, cap=%d\n", slice, len(slice), cap(slice))

	/*
		Visual representation:
		Underlying array: [0, 0, 0, _, _]
		                   ^     ^     ^
		                   |     |     |
		                 start  len   cap
	*/

	// Visualize length vs capacity with array slicing
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice1 := array[2:5] // Start at index 2, end before index 5
	fmt.Printf("\narray[2:5]: %v | len=%d, cap=%d\n", slice1, len(slice1), cap(slice1))

	slice2 := array[2:5:7] // Three-index slice: [low:high:max]
	fmt.Printf("array[2:5:7]: %v | len=%d, cap=%d\n", slice2, len(slice2), cap(slice2))

	// Capacity determines how much slice can grow without reallocation
	fmt.Println("\n📊 Capacity Growth Demo:")
	growth := make([]int, 0)
	prevCap := 0
	for i := 0; i < 20; i++ {
		growth = append(growth, i)
		if cap(growth) != prevCap {
			fmt.Printf("   len=%2d, cap changed: %d → %d\n", len(growth), prevCap, cap(growth))
			prevCap = cap(growth)
		}
	}
}

// ============================================================
// 4️⃣ SLICING OPERATIONS (Sub-slices)
// ============================================================
func SlicingOperations() {
	fmt.Println("\n📚 4. SLICING OPERATIONS (Sub-slices)")
	fmt.Println("─" + "───────────────────────────────────────────────")

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Original:", slice)

	/*
		Slice syntax: slice[low:high]
		- low: starting index (inclusive)
		- high: ending index (exclusive)
		- Result length: high - low
	*/

	// Basic slicing
	fmt.Println("\n🔹 Basic Slicing:")
	fmt.Println("slice[2:5]  →", slice[2:5]) // Elements at index 2, 3, 4
	fmt.Println("slice[:5]   →", slice[:5])  // First 5 elements (0 to 4)
	fmt.Println("slice[5:]   →", slice[5:])  // From index 5 to end
	fmt.Println("slice[:]    →", slice[:])   // Entire slice (copy reference)

	// Negative thinking - there are NO negative indices in Go!
	// slice[-1] will cause compile error

	// Getting last N elements
	fmt.Println("\n🔹 Last N Elements:")
	fmt.Println("Last 3:", slice[len(slice)-3:])
	fmt.Println("Last 1:", slice[len(slice)-1:])

	// Getting first N elements
	fmt.Println("\n🔹 First N Elements:")
	fmt.Println("First 3:", slice[:3])

	// Remove first element
	fmt.Println("\n🔹 Remove Elements:")
	fmt.Println("Remove first:", slice[1:])

	// Remove last element
	fmt.Println("Remove last:", slice[:len(slice)-1])

	// Three-index slicing [low:high:max]
	fmt.Println("\n🔹 Three-Index Slicing:")
	slice2 := slice[2:5:6] // cap = max - low = 6 - 2 = 4
	fmt.Printf("slice[2:5:6]: %v | len=%d, cap=%d\n", slice2, len(slice2), cap(slice2))

	// ⚠️ Important: Slices share underlying array!
	fmt.Println("\n⚠️  Shared Memory Demo:")
	original := []int{1, 2, 3, 4, 5}
	subSlice := original[1:4]
	fmt.Println("Original:", original)
	fmt.Println("SubSlice:", subSlice)
	subSlice[0] = 999 // This modifies original too!
	fmt.Println("After modifying subSlice[0]:")
	fmt.Println("Original:", original)
	fmt.Println("SubSlice:", subSlice)
}

// ============================================================
// 5️⃣ APPENDING TO SLICES
// ============================================================
func AppendingToSlices() {
	fmt.Println("\n📚 5. APPENDING TO SLICES")
	fmt.Println("─" + "───────────────────────────────────────────────")

	// append() is a built-in function that adds elements to a slice
	// IMPORTANT: append may return a NEW slice if capacity is exceeded!

	slice := []int{1, 2, 3}
	fmt.Printf("Initial: %v | len=%d, cap=%d\n", slice, len(slice), cap(slice))

	// Append single element
	slice = append(slice, 4)
	fmt.Printf("After append(4): %v | len=%d, cap=%d\n", slice, len(slice), cap(slice))

	// Append multiple elements
	slice = append(slice, 5, 6, 7)
	fmt.Printf("After append(5,6,7): %v | len=%d, cap=%d\n", slice, len(slice), cap(slice))

	// Append another slice using ... operator
	more := []int{8, 9, 10}
	slice = append(slice, more...)
	fmt.Printf("After append(more...): %v | len=%d, cap=%d\n", slice, len(slice), cap(slice))

	// ⚠️ Common Mistake: Forgetting to assign append result
	fmt.Println("\n⚠️  Common Mistake Demo:")
	test := []int{1, 2}
	fmt.Println("Before (cap 2):", test)
	_ = append(test, 3) // WRONG! Result discarded, test unchanged
	fmt.Println("After wrong append:", test)
	test = append(test, 3) // CORRECT! Assign the result
	fmt.Println("After correct append:", test)

	// Insert element at index
	fmt.Println("\n🔹 Insert at Index:")
	insertAt := func(slice []int, index int, value int) []int {
		slice = append(slice[:index+1], slice[index:]...)
		slice[index] = value
		return slice
	}
	data := []int{1, 2, 4, 5}
	data = insertAt(data, 2, 3)
	fmt.Println("Insert 3 at index 2:", data)

	// Delete element at index
	fmt.Println("\n🔹 Delete at Index:")
	deleteAt := func(slice []int, index int) []int {
		return append(slice[:index], slice[index+1:]...)
	}
	data = deleteAt(data, 2)
	fmt.Println("Delete at index 2:", data)

	// Prepend element
	fmt.Println("\n🔹 Prepend Element:")
	prepend := append([]int{0}, data...)
	fmt.Println("Prepend 0:", prepend)
}

// ============================================================
// 6️⃣ COPYING SLICES
// ============================================================
func CopyingSlices() {
	fmt.Println("\n📚 6. COPYING SLICES")
	fmt.Println("─" + "───────────────────────────────────────────────")

	// copy() is a built-in function that copies elements between slices

	// Basic copy
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copied := copy(dst, src)
	fmt.Printf("Copied %d elements: %v\n", copied, dst)

	// dst must have enough space
	smallDst := make([]int, 3)
	copied = copy(smallDst, src)
	fmt.Printf("Copy to smaller slice (%d copied): %v\n", copied, smallDst)

	// Copy part of slice
	partialDst := make([]int, 5)
	copy(partialDst[2:], src[:3]) // Copy first 3 of src to last 3 of dst
	fmt.Println("Partial copy:", partialDst)

	// Deep copy vs shallow copy for nested slices
	fmt.Println("\n🔹 Deep Copy for 2D Slices:")
	matrix := [][]int{{1, 2}, {3, 4}}
	shallowCopy := make([][]int, len(matrix))
	copy(shallowCopy, matrix) // Only copies the outer slice!
	shallowCopy[0][0] = 999
	fmt.Println("Original after shallow copy modification:", matrix)

	// Proper deep copy
	matrix2 := [][]int{{1, 2}, {3, 4}}
	deepCopy := make([][]int, len(matrix2))
	for i := range matrix2 {
		deepCopy[i] = make([]int, len(matrix2[i]))
		copy(deepCopy[i], matrix2[i])
	}
	deepCopy[0][0] = 999
	fmt.Println("Original after deep copy modification:", matrix2)

	// Alternative: Clone using slices.Clone (Go 1.21+)
	original := []int{1, 2, 3, 4, 5}
	cloned := slices.Clone(original)
	fmt.Println("Cloned slice:", cloned)
}

// ============================================================
// 7️⃣ NIL VS EMPTY SLICE
// ============================================================
func NilVsEmptySlice() {
	fmt.Println("\n📚 7. NIL VS EMPTY SLICE")
	fmt.Println("─" + "───────────────────────────────────────────────")

	// Nil slice: no underlying array
	var nilSlice []int

	// Empty slice: has underlying array with 0 capacity
	emptySlice1 := []int{}
	emptySlice2 := make([]int, 0)

	fmt.Println("🔹 Comparison:")
	fmt.Printf("nilSlice:    %v | len=%d, cap=%d, nil=%t\n",
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("emptySlice1: %v | len=%d, cap=%d, nil=%t\n",
		emptySlice1, len(emptySlice1), cap(emptySlice1), emptySlice1 == nil)
	fmt.Printf("emptySlice2: %v | len=%d, cap=%d, nil=%t\n",
		emptySlice2, len(emptySlice2), cap(emptySlice2), emptySlice2 == nil)

	// They're functionally equivalent for most operations!
	fmt.Println("\n🔹 Both work the same way:")
	nilSlice = append(nilSlice, 1)
	emptySlice1 = append(emptySlice1, 1)
	fmt.Println("After append - nilSlice:", nilSlice)
	fmt.Println("After append - emptySlice:", emptySlice1)

	// JSON encoding difference
	fmt.Println("\n🔹 JSON Encoding Difference:")
	fmt.Println("nil slice would encode to: null")
	fmt.Println("empty slice would encode to: []")

	// Best practice: prefer nil slice declaration
	fmt.Println("\n✅ Best Practice:")
	fmt.Println("Use 'var s []int' for nil slice")
	fmt.Println("Use 'make([]int, 0)' or '[]int{}' when you specifically need non-nil")
}

// ============================================================
// 8️⃣ MULTI-DIMENSIONAL SLICES
// ============================================================
func MultiDimensionalSlices() {
	fmt.Println("\n📚 8. MULTI-DIMENSIONAL SLICES")
	fmt.Println("─" + "───────────────────────────────────────────────")

	// 2D slice (slice of slices)
	fmt.Println("🔹 2D Slice:")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:")
	for i, row := range matrix {
		fmt.Printf("  Row %d: %v\n", i, row)
	}

	// Creating 2D slice with make
	fmt.Println("\n🔹 Creating 2D Slice with make:")
	rows, cols := 3, 4
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}
	// Fill with values
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			grid[i][j] = i*cols + j
		}
	}
	fmt.Println("Grid:")
	for i, row := range grid {
		fmt.Printf("  Row %d: %v\n", i, row)
	}

	// Jagged slices (rows with different lengths)
	fmt.Println("\n🔹 Jagged Slice (Triangle):")
	jagged := [][]int{
		{1},
		{2, 3},
		{4, 5, 6},
		{7, 8, 9, 10},
	}
	for i, row := range jagged {
		fmt.Printf("  Row %d: %v\n", i, row)
	}

	// 3D slice
	fmt.Println("\n🔹 3D Slice:")
	cube := [][][]int{
		{{1, 2}, {3, 4}},
		{{5, 6}, {7, 8}},
	}
	fmt.Printf("Cube[0][1][0] = %d\n", cube[0][1][0])
}

// ============================================================
// 9️⃣ SLICE ITERATION
// ============================================================
func SliceIteration() {
	fmt.Println("\n📚 9. SLICE ITERATION")
	fmt.Println("─" + "───────────────────────────────────────────────")

	slice := []string{"Go", "is", "awesome", "!"}

	// Method 1: for loop with index
	fmt.Println("🔹 Method 1: for loop with index")
	for i := 0; i < len(slice); i++ {
		fmt.Printf("  slice[%d] = %s\n", i, slice[i])
	}

	// Method 2: range with index and value
	fmt.Println("\n🔹 Method 2: range with index and value")
	for index, value := range slice {
		fmt.Printf("  Index %d: %s\n", index, value)
	}

	// Method 3: range with only index
	fmt.Println("\n🔹 Method 3: range with only index")
	for i := range slice {
		fmt.Printf("  Index: %d\n", i)
	}

	// Method 4: range with only value (use blank identifier for index)
	fmt.Println("\n🔹 Method 4: range with only value")
	for _, value := range slice {
		fmt.Printf("  Value: %s\n", value)
	}

	// ⚠️ Important: range creates a COPY of each element
	fmt.Println("\n⚠️  Range Copy Gotcha:")
	nums := []int{1, 2, 3}
	for _, v := range nums {
		v = v * 2 // This doesn't modify the slice!
	}
	fmt.Println("After range (wrong):", nums)

	// Correct way to modify during iteration
	for i := range nums {
		nums[i] = nums[i] * 2
	}
	fmt.Println("After range (correct):", nums)

	// Reverse iteration
	fmt.Println("\n🔹 Reverse Iteration:")
	for i := len(slice) - 1; i >= 0; i-- {
		fmt.Printf("  slice[%d] = %s\n", i, slice[i])
	}
}

// ============================================================
// 🔟 SLICE SORTING
// ============================================================
func SliceSorting() {
	fmt.Println("\n📚 10. SLICE SORTING")
	fmt.Println("─" + "───────────────────────────────────────────────")

	// Sort integers
	fmt.Println("🔹 Sort Integers:")
	ints := []int{5, 2, 6, 3, 1, 4}
	fmt.Println("Before:", ints)
	sort.Ints(ints)
	fmt.Println("After sort.Ints:", ints)

	// Sort strings
	fmt.Println("\n🔹 Sort Strings:")
	strings := []string{"banana", "apple", "cherry", "date"}
	fmt.Println("Before:", strings)
	sort.Strings(strings)
	fmt.Println("After sort.Strings:", strings)

	// Sort floats
	fmt.Println("\n🔹 Sort Float64s:")
	floats := []float64{3.14, 1.41, 2.71, 1.73}
	fmt.Println("Before:", floats)
	sort.Float64s(floats)
	fmt.Println("After sort.Float64s:", floats)

	// Sort in reverse (descending)
	fmt.Println("\n🔹 Sort Descending:")
	nums := []int{5, 2, 6, 3, 1, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println("Descending:", nums)

	// Custom sort with sort.Slice
	fmt.Println("\n🔹 Custom Sort:")
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	fmt.Println("Before:", people)

	// Sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("After sort by age:", people)

	// Check if sorted
	fmt.Println("\n🔹 Check if Sorted:")
	sorted := []int{1, 2, 3, 4, 5}
	unsorted := []int{5, 2, 3, 1, 4}
	fmt.Printf("Is %v sorted? %t\n", sorted, sort.IntsAreSorted(sorted))
	fmt.Printf("Is %v sorted? %t\n", unsorted, sort.IntsAreSorted(unsorted))

	// Stable sort (maintains original order of equal elements)
	fmt.Println("\n🔹 Stable Sort:")
	stablePeople := []Person{
		{"Alice", 30},
		{"Bob", 30},
		{"Charlie", 25},
	}
	sort.SliceStable(stablePeople, func(i, j int) bool {
		return stablePeople[i].Age < stablePeople[j].Age
	})
	fmt.Println("Stable sorted:", stablePeople)
}

// ============================================================
// 1️⃣1️⃣ SLICE SEARCHING
// ============================================================
func SliceSearching() {
	fmt.Println("\n📚 11. SLICE SEARCHING")
	fmt.Println("─" + "───────────────────────────────────────────────")

	// Linear search (find index)
	fmt.Println("🔹 Linear Search:")
	slice := []int{10, 20, 30, 40, 50}
	target := 30

	indexOf := func(s []int, target int) int {
		for i, v := range s {
			if v == target {
				return i
			}
		}
		return -1
	}
	fmt.Printf("Index of %d in %v: %d\n", target, slice, indexOf(slice, target))

	// Contains check
	fmt.Println("\n🔹 Contains Check:")
	contains := func(s []int, target int) bool {
		for _, v := range s {
			if v == target {
				return true
			}
		}
		return false
	}
	fmt.Printf("Contains 30? %t\n", contains(slice, 30))
	fmt.Printf("Contains 99? %t\n", contains(slice, 99))

	// Using slices.Contains (Go 1.21+)
	fmt.Printf("slices.Contains(30)? %t\n", slices.Contains(slice, 30))
	fmt.Printf("slices.Contains(99)? %t\n", slices.Contains(slice, 99))

	// Binary search (requires sorted slice)
	fmt.Println("\n🔹 Binary Search (sorted slice required):")
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	searchVal := 7
	index := sort.SearchInts(sorted, searchVal)
	if index < len(sorted) && sorted[index] == searchVal {
		fmt.Printf("Found %d at index %d\n", searchVal, index)
	} else {
		fmt.Printf("%d not found\n", searchVal)
	}

	// Find index using slices.Index (Go 1.21+)
	fmt.Println("\n🔹 Using slices.Index (Go 1.21+):")
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Printf("Index of 'banana': %d\n", slices.Index(fruits, "banana"))
	fmt.Printf("Index of 'grape': %d\n", slices.Index(fruits, "grape"))

	// Find with custom condition
	fmt.Println("\n🔹 Find with Custom Condition:")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	firstEvenIndex := slices.IndexFunc(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("First even number at index: %d (value: %d)\n", firstEvenIndex, nums[firstEvenIndex])
}

// ============================================================
// 1️⃣2️⃣ SLICE FILTERING
// ============================================================
func SliceFiltering() {
	fmt.Println("\n📚 12. SLICE FILTERING")
	fmt.Println("─" + "───────────────────────────────────────────────")

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original:", nums)

	// Filter even numbers
	fmt.Println("\n🔹 Filter Even Numbers:")
	filter := func(s []int, predicate func(int) bool) []int {
		result := make([]int, 0)
		for _, v := range s {
			if predicate(v) {
				result = append(result, v)
			}
		}
		return result
	}

	evens := filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("Even numbers:", evens)

	odds := filter(nums, func(n int) bool { return n%2 != 0 })
	fmt.Println("Odd numbers:", odds)

	greaterThan5 := filter(nums, func(n int) bool { return n > 5 })
	fmt.Println("Greater than 5:", greaterThan5)

	// Filter in-place (modifies original slice)
	fmt.Println("\n🔹 Filter In-Place (Efficient):")
	filterInPlace := func(s []int, predicate func(int) bool) []int {
		n := 0
		for _, v := range s {
			if predicate(v) {
				s[n] = v
				n++
			}
		}
		return s[:n]
	}

	numsCopy := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numsCopy = filterInPlace(numsCopy, func(n int) bool { return n%2 == 0 })
	fmt.Println("Filtered in-place:", numsCopy)

	// Using slices.DeleteFunc (Go 1.21+)
	fmt.Println("\n🔹 Using slices.DeleteFunc (Go 1.21+):")
	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Delete odd numbers (keep even)
	nums2 = slices.DeleteFunc(nums2, func(n int) bool { return n%2 != 0 })
	fmt.Println("After DeleteFunc (remove odds):", nums2)
}

// ============================================================
// 1️⃣3️⃣ SLICE MAPPING (Transform)
// ============================================================
func SliceMapping() {
	fmt.Println("\n📚 13. SLICE MAPPING (Transform)")
	fmt.Println("─" + "───────────────────────────────────────────────")

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Original:", nums)

	// Map: Transform each element
	fmt.Println("\n🔹 Map Function:")
	mapSlice := func(s []int, f func(int) int) []int {
		result := make([]int, len(s))
		for i, v := range s {
			result[i] = f(v)
		}
		return result
	}

	doubled := mapSlice(nums, func(n int) int { return n * 2 })
	fmt.Println("Doubled:", doubled)

	squared := mapSlice(nums, func(n int) int { return n * n })
	fmt.Println("Squared:", squared)

	// Reduce: Combine all elements into single value
	fmt.Println("\n🔹 Reduce Function:")
	reduce := func(s []int, initial int, f func(int, int) int) int {
		result := initial
		for _, v := range s {
			result = f(result, v)
		}
		return result
	}

	sum := reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Println("Sum:", sum)

	product := reduce(nums, 1, func(acc, n int) int { return acc * n })
	fmt.Println("Product:", product)

	max := reduce(nums, nums[0], func(acc, n int) int {
		if n > acc {
			return n
		}
		return acc
	})
	fmt.Println("Max:", max)

	// FlatMap: Map and flatten
	fmt.Println("\n🔹 FlatMap:")
	flatMap := func(s []int, f func(int) []int) []int {
		result := make([]int, 0)
		for _, v := range s {
			result = append(result, f(v)...)
		}
		return result
	}

	duplicated := flatMap(nums, func(n int) []int { return []int{n, n} })
	fmt.Println("Duplicated:", duplicated)

	// Chunk: Split slice into chunks
	fmt.Println("\n🔹 Chunk:")
	chunk := func(s []int, size int) [][]int {
		var chunks [][]int
		for size < len(s) {
			s, chunks = s[size:], append(chunks, s[0:size:size])
		}
		chunks = append(chunks, s)
		return chunks
	}

	chunked := chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
	fmt.Println("Chunked by 3:", chunked)
}

// ============================================================
// 1️⃣4️⃣ SLICE COMPARISON
// ============================================================
func SliceComparison() {
	fmt.Println("\n📚 14. SLICE COMPARISON")
	fmt.Println("─" + "───────────────────────────────────────────────")

	// ⚠️ You cannot compare slices with == operator!
	// slice1 == slice2 // This won't compile!

	// Method 1: reflect.DeepEqual (slower, works for any type)
	fmt.Println("🔹 Using reflect.DeepEqual:")
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slice3 := []int{1, 2, 4}

	fmt.Printf("%v == %v ? %t\n", slice1, slice2, reflect.DeepEqual(slice1, slice2))
	fmt.Printf("%v == %v ? %t\n", slice1, slice3, reflect.DeepEqual(slice1, slice3))

	// Method 2: slices.Equal (Go 1.21+, faster)
	fmt.Println("\n🔹 Using slices.Equal (Go 1.21+):")
	fmt.Printf("%v == %v ? %t\n", slice1, slice2, slices.Equal(slice1, slice2))
	fmt.Printf("%v == %v ? %t\n", slice1, slice3, slices.Equal(slice1, slice3))

	// Method 3: Manual comparison
	fmt.Println("\n🔹 Manual Comparison:")
	equalSlices := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	fmt.Printf("%v == %v ? %t\n", slice1, slice2, equalSlices(slice1, slice2))

	// Compare with custom function
	fmt.Println("\n🔹 Custom Comparison:")
	type Item struct {
		ID   int
		Name string
	}
	items1 := []Item{{1, "a"}, {2, "b"}}
	items2 := []Item{{1, "a"}, {2, "b"}}

	equalItems := slices.EqualFunc(items1, items2, func(a, b Item) bool {
		return a.ID == b.ID && a.Name == b.Name
	})
	fmt.Printf("Custom struct comparison: %t\n", equalItems)

	// Compare slice to nil
	fmt.Println("\n🔹 Nil Comparison:")
	var nilSlice []int
	emptySlice := []int{}
	fmt.Printf("nil slice == nil? %t\n", nilSlice == nil)
	fmt.Printf("empty slice == nil? %t\n", emptySlice == nil)
}

// ============================================================
// 1️⃣5️⃣ SLICE MEMORY INTERNALS
// ============================================================
func SliceMemoryInternals() {
	fmt.Println("\n📚 15. SLICE MEMORY INTERNALS")
	fmt.Println("─" + "───────────────────────────────────────────────")

	/*
		Slice header structure:
		type sliceHeader struct {
			Data uintptr  // pointer to underlying array
			Len  int      // number of elements
			Cap  int      // capacity
		}
	*/

	fmt.Println("🔹 Slice Header Size:")
	var s []int
	fmt.Printf("Size of slice header: %d bytes\n", 24) // 8 + 8 + 8 on 64-bit

	// Demonstrate underlying array sharing
	fmt.Println("\n🔹 Underlying Array Sharing:")
	original := []int{1, 2, 3, 4, 5}
	slice1 := original[1:4]
	slice2 := original[2:5]

	fmt.Printf("original: %v (ptr: %p)\n", original, original)
	fmt.Printf("slice1[1:4]: %v (ptr: %p)\n", slice1, slice1)
	fmt.Printf("slice2[2:5]: %v (ptr: %p)\n", slice2, slice2)

	// They share memory!
	slice1[1] = 999
	fmt.Println("\nAfter slice1[1] = 999:")
	fmt.Println("original:", original)
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)

	// When does reallocation happen?
	fmt.Println("\n🔹 Reallocation Behavior:")
	s = make([]int, 0, 4)
	fmt.Printf("Initial capacity: %d\n", cap(s))

	for i := 0; i < 10; i++ {
		prevCap := cap(s)
		s = append(s, i)
		if cap(s) != prevCap {
			fmt.Printf("Grew from cap %d to %d at len %d\n", prevCap, cap(s), len(s))
		}
	}

	// Memory growth pattern
	fmt.Println("\n🔹 Go's Growth Strategy:")
	fmt.Println("- For small slices: doubles capacity")
	fmt.Println("- For large slices (>1024): grows by ~25%")

	// Prevent memory leaks with large slices
	fmt.Println("\n🔹 Memory Leak Prevention:")
	fmt.Println("When taking small slice from large slice, copy to avoid retaining large array:")
	largeSlice := make([]int, 1000000)
	largeSlice[0] = 42

	// BAD: retains all 1M elements in memory
	// smallRef := largeSlice[:1]

	// GOOD: only keeps needed elements
	smallCopy := make([]int, 1)
	copy(smallCopy, largeSlice[:1])
	fmt.Printf("Small copy: %v (backed by tiny array)\n", smallCopy)
}

// ============================================================
// 1️⃣6️⃣ COMMON GOTCHAS
// ============================================================
func SliceGotchas() {
	fmt.Println("\n📚 16. COMMON GOTCHAS")
	fmt.Println("─" + "───────────────────────────────────────────────")

	// Gotcha 1: Append may or may not modify original
	fmt.Println("⚠️  Gotcha 1: Append Behavior")
	a := make([]int, 3, 6) // len=3, cap=6
	a[0], a[1], a[2] = 1, 2, 3
	b := append(a, 4) // Fits in capacity, modifies underlying array
	fmt.Printf("a: %v, b: %v\n", a, b)
	b[0] = 999
	fmt.Printf("After b[0]=999: a=%v (a is affected because same array!)\n", a)

	// Gotcha 2: Loop variable capture
	fmt.Println("\n⚠️  Gotcha 2: Loop Variable Capture")
	funcs := make([]func(), 3)
	nums := []int{1, 2, 3}

	// WRONG (before Go 1.22)
	// for i, n := range nums {
	//     funcs[i] = func() { fmt.Println(n) } // All print 3!
	// }

	// CORRECT (or default in Go 1.22+)
	for i, n := range nums {
		n := n // Capture in local scope
		funcs[i] = func() { fmt.Println(n) }
	}
	fmt.Print("Loop capture results: ")
	for _, f := range funcs {
		f()
	}

	// Gotcha 3: Range returns copy
	fmt.Println("\n⚠️  Gotcha 3: Range Returns Copy")
	type Point struct{ X, Y int }
	points := []Point{{1, 2}, {3, 4}}
	for _, p := range points {
		p.X = 100 // Modifying copy, not original!
	}
	fmt.Println("After range modification:", points) // Unchanged!

	// Correct way:
	for i := range points {
		points[i].X = 100
	}
	fmt.Println("After index modification:", points) // Changed!

	// Gotcha 4: Nil slice vs empty slice in JSON
	fmt.Println("\n⚠️  Gotcha 4: Nil vs Empty in JSON")
	fmt.Println("var s []int  → JSON: null")
	fmt.Println("s := []int{} → JSON: []")

	// Gotcha 5: Deleting while iterating
	fmt.Println("\n⚠️  Gotcha 5: Deleting While Iterating")
	fmt.Println("Never delete elements while iterating forward!")
	fmt.Println("Iterate backward or create new slice instead.")
}

// ============================================================
// 1️⃣7️⃣ USEFUL TRICKS AND PATTERNS
// ============================================================
func SliceTricksAndPatterns() {
	fmt.Println("\n📚 17. USEFUL TRICKS AND PATTERNS")
	fmt.Println("─" + "───────────────────────────────────────────────")

	// Trick 1: Reverse a slice
	fmt.Println("🔸 Reverse Slice:")
	reverse := func(s []int) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}
	nums := []int{1, 2, 3, 4, 5}
	reverse(nums)
	fmt.Println("Reversed:", nums)

	// Trick 2: Remove duplicates
	fmt.Println("\n🔸 Remove Duplicates:")
	removeDuplicates := func(s []int) []int {
		seen := make(map[int]bool)
		result := make([]int, 0)
		for _, v := range s {
			if !seen[v] {
				seen[v] = true
				result = append(result, v)
			}
		}
		return result
	}
	withDups := []int{1, 2, 2, 3, 3, 3, 4}
	fmt.Println("With duplicates:", withDups)
	fmt.Println("Without duplicates:", removeDuplicates(withDups))

	// Trick 3: Shuffle
	fmt.Println("\n🔸 Shuffle Slice:")
	// Note: In production, use math/rand/v2 or crypto/rand
	shuffled := []int{1, 2, 3, 4, 5}
	fmt.Println("Before shuffle:", shuffled)
	// Simple demonstration (for real shuffle, use rand.Shuffle)

	// Trick 4: Pop from slice
	fmt.Println("\n🔸 Stack Operations:")
	stack := []int{1, 2, 3}
	fmt.Println("Stack:", stack)

	// Push
	stack = append(stack, 4)
	fmt.Println("After push(4):", stack)

	// Pop
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Printf("Pop: %d, Stack: %v\n", top, stack)

	// Trick 5: Queue operations
	fmt.Println("\n🔸 Queue Operations:")
	queue := []int{1, 2, 3}
	fmt.Println("Queue:", queue)

	// Enqueue
	queue = append(queue, 4)
	fmt.Println("After enqueue(4):", queue)

	// Dequeue
	front := queue[0]
	queue = queue[1:]
	fmt.Printf("Dequeue: %d, Queue: %v\n", front, queue)

	// Trick 6: Rotate slice
	fmt.Println("\n🔸 Rotate Slice:")
	rotate := func(s []int, k int) []int {
		k = k % len(s)
		return append(s[k:], s[:k]...)
	}
	original := []int{1, 2, 3, 4, 5}
	rotated := rotate(original, 2)
	fmt.Printf("Rotate %v by 2: %v\n", original, rotated)

	// Trick 7: Flatten nested slice
	fmt.Println("\n🔸 Flatten 2D Slice:")
	flatten := func(s [][]int) []int {
		result := make([]int, 0)
		for _, row := range s {
			result = append(result, row...)
		}
		return result
	}
	nested := [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println("Flattened:", flatten(nested))

	// Trick 8: Zip two slices
	fmt.Println("\n🔸 Zip Slices:")
	zip := func(a, b []int) [][]int {
		minLen := len(a)
		if len(b) < minLen {
			minLen = len(b)
		}
		result := make([][]int, minLen)
		for i := 0; i < minLen; i++ {
			result[i] = []int{a[i], b[i]}
		}
		return result
	}
	zipped := zip([]int{1, 2, 3}, []int{4, 5, 6})
	fmt.Println("Zipped:", zipped)
}

// ============================================================
// 1️⃣8️⃣ MODERN SLICES PACKAGE (Go 1.21+)
// ============================================================
func ModernSlicePackage() {
	fmt.Println("\n📚 18. MODERN SLICES PACKAGE (Go 1.21+)")
	fmt.Println("─" + "───────────────────────────────────────────────")

	fmt.Println("Import: \"slices\"")
	fmt.Println()

	// slices.Sort
	fmt.Println("🔸 slices.Sort:")
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}
	slices.Sort(nums)
	fmt.Println("Sorted:", nums)

	// slices.BinarySearch
	fmt.Println("\n🔸 slices.BinarySearch:")
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	idx, found := slices.BinarySearch(sorted, 5)
	fmt.Printf("BinarySearch(5): index=%d, found=%t\n", idx, found)

	// slices.Contains
	fmt.Println("\n🔸 slices.Contains:")
	fmt.Printf("Contains 5? %t\n", slices.Contains(sorted, 5))
	fmt.Printf("Contains 99? %t\n", slices.Contains(sorted, 99))

	// slices.Index
	fmt.Println("\n🔸 slices.Index:")
	fmt.Printf("Index of 5: %d\n", slices.Index(sorted, 5))

	// slices.Min / slices.Max
	fmt.Println("\n🔸 slices.Min / slices.Max:")
	values := []int{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Printf("Min: %d, Max: %d\n", slices.Min(values), slices.Max(values))

	// slices.Reverse
	fmt.Println("\n🔸 slices.Reverse:")
	toReverse := []int{1, 2, 3, 4, 5}
	slices.Reverse(toReverse)
	fmt.Println("Reversed:", toReverse)

	// slices.Clone
	fmt.Println("\n🔸 slices.Clone:")
	original := []int{1, 2, 3}
	cloned := slices.Clone(original)
	fmt.Println("Cloned:", cloned)

	// slices.Compact (remove consecutive duplicates)
	fmt.Println("\n🔸 slices.Compact:")
	withDups := []int{1, 1, 2, 2, 2, 3, 3}
	compacted := slices.Compact(withDups)
	fmt.Println("Compacted:", compacted)

	// slices.Equal
	fmt.Println("\n🔸 slices.Equal:")
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 4}
	fmt.Printf("%v == %v ? %t\n", a, b, slices.Equal(a, b))
	fmt.Printf("%v == %v ? %t\n", a, c, slices.Equal(a, c))

	// slices.Insert
	fmt.Println("\n🔸 slices.Insert:")
	s := []int{1, 2, 5}
	s = slices.Insert(s, 2, 3, 4)
	fmt.Println("After Insert(2, 3, 4):", s)

	// slices.Delete
	fmt.Println("\n🔸 slices.Delete:")
	s = []int{1, 2, 3, 4, 5}
	s = slices.Delete(s, 2, 4) // Delete elements from index 2 to 4
	fmt.Println("After Delete(2, 4):", s)

	// slices.Replace
	fmt.Println("\n🔸 slices.Replace:")
	s = []int{1, 2, 3, 4, 5}
	s = slices.Replace(s, 1, 4, 10, 20)
	fmt.Println("After Replace(1, 4, 10, 20):", s)

	// slices.IsSorted
	fmt.Println("\n🔸 slices.IsSorted:")
	fmt.Printf("Is [1,2,3,4,5] sorted? %t\n", slices.IsSorted([]int{1, 2, 3, 4, 5}))
	fmt.Printf("Is [5,2,3,1,4] sorted? %t\n", slices.IsSorted([]int{5, 2, 3, 1, 4}))

	fmt.Println("\n" + "=" + "=================================================")
	fmt.Println("🎉 END OF SLICE TUTORIAL")
	fmt.Println("=" + "=================================================")
}
