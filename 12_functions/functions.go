package main

// func add(a int, b int) int {
// 	return a + b
// }

// func add(a, b int) int {   --if both are int then we can write like this
// 	return a + b
// }

// func getLanguages() (string, string, bool) {
// 	return "golang", "javascript", true
// }

// func processIt(fn func(a int) int) {
// 	fn(1)
// }

// func processIt() func(a int) int {
// 	return func(a int) int {
// 		return 4
// 	}
// }

func main() {
	// fn := func(a int) int {
	// 	return 2
	// }

	fn := processIt()
	fn(6)
   // fmt.Printlm(getLanguages()) 
	//let's first get and then print(below code)
	// lang1, lang2, _ := getLanguages()    --if we don't want to use lang3 then we are putting underscore(_)
	// fmt.Println(lang1, lang2)
	// result := add(3, 5)
	// fmt.Println(result)
}
