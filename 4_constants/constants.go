package main

import "fmt"

// const age = 30

func main() {
	// :=
	// const name = "golang"
	// const age = 30

	// fmt.Println(age)

	const (               --if we are printing list of const we can use const like this
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)

}
