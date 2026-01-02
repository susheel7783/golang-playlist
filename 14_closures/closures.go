package main

import "fmt"

func counter() func() int {   --here counter function not receiving any parameter and it returns an integer type function
	var count int = 0

	return func() int { -- here we can make any anonymous function, actually this will be our increment function  
		count += 1
		return count
	}

}

func main() {
	increment := counter()    -- to call increment function we have to call counter function and this will return the increment function

	fmt.Println(increment())  --op 1
	fmt.Println(increment())  --op 2  
}
