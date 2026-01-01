package main

// maps -> hash, object, dict
func main() {
	// creating map

	// m := make(map[string]string)  --key- string, value- string 

	// setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"

	// get an element
	// fmt.Println(m["name"], m["area"])
	// IMP: if key does not exists in the map then it returns zero value

	// m := make(map[string]int) --key- string, value- int
	// m["age"] = 30
	// m["price"] = 50
	// fmt.Println(m["phone"])
	// fmt.Println(len(m))

	// delete(m, "price")
	// clear(m)  -- we can empty the map

	// fmt.Println(m)
	// fmt.Println(m)

	// m := map[string]int{"price": 40, "phones": 3}  -- we can make map without using make fun

	// v, ok := m["phones"]    --first var give value and 2nd give true or  false
	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	// m1 := map[string]int{"price": 40, "phones": 3}
	// m2 := map[string]int{"price": 40, "phones": 8}
	// fmt.Println(maps.Equal(m1, m2))

}
