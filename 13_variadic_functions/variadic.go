package main

import "fmt"

func sum(nums ...int) int {    --here we are passing n numbers of parameters using (...) and passing all the parameter are int
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func main() {
    fmt.Println(1,2,3,4,5,66, "Hi")     --we can pass anything here, this is a variadic function, we can pass n numbers of parameters
	result:=sum(1,2,3,6,7)  --here  we can pass any number of parameters 
	
	nums := []int{3, 4, 5, 6}
	result := sum(nums...)
	fmt.Println(result)
}
