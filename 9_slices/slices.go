package main

// slice -> dynamic array (expand the size)
//this is most used construct in go
// + useful methods
func main() {
	// uninitialized slice is nil   --in go we write nil  instead of null
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)   --if we don't want nil we can initialize like this
	// capacity -> maximum numbers of elements can fit and it will resize itself
	// fmt.Println(cap(nums)) --cap function gives the capacity
	// fmt.Println(nums == nil)

	// nums := []int{}   --we can create slice like this as well

	// nums = append(nums, 1)  -- we can use append method
	// nums = append(nums, 2)

	
    // var nums = make([]int, 2, 5)   --here 2 element 5 capacity we initilaise 
	// nums[0] = 3
	// nums[1] = 5
	// fmt.Println(nums)   --op [3,5]
	// fmt.Println(cap(nums)) --op 5
	// fmt.Println(len(nums)) --op 2
	

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))
    // fmt.Println(nums, nums2)   --op [2] []

	
	// // copy function
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)  --op [2] [2]

	// slice operator

	// var nums = []int{1, 2, 3, 4, 5}
	// fmt.Println(nums[0:1])   --op [1]
	// fmt.Println(nums[:2])    --op [1,2]
	// fmt.Println(nums[1:])    --op [2,3,4,5]

	// slices
	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 4}

	// fmt.Println(slices.Equal(nums1, nums2))  --op false

	// var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	// fmt.Println(nums)

}
