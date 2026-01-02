package main

import "fmt"

// iterating over data structures
func main() {
	// nums := []int{6, 7, 8}
    // for i:=0;i<len(nums);i++{
   // 	fmt.Println(nums[i])  --op 6 7 8 
   // }


	//we will iterate using range

	// for _, num := range nums {
	// 	fmt.Println(num)           --op 6 7 8
	// }

	
	// for i, num := range nums {    -- i is basically Index 
	// 	fmt.Println(num, i)
	// }

	// sum:=0
	// for _, num := range nums {
	// 	sum=sum+num
	// }
   // fmt.Println(sum)         --21

	
	// m := map[string]string{"fname": "john", "lname": "doe"}

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// 	fmt.Println(k, v)
	// }

	// for k := range m {    --it will give only keys
	// 	fmt.Println(k)
	// }

	// unicode code point rune
	// starting byte of rune
	// 300 -> 1 byte , 2 byte
// for i, c := range "golang" {
// 		fmt.Println(i, c)             --this will give unicode of the cahr
// 	}
	
	for i, c := range "golang" {
		fmt.Println(i, string(c))   --this will give char
	}

}
