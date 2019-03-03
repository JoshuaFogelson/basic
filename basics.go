package main

import( 
	"fmt"
	"errors"
	"math"
)

type person struct {
	name string
	age int
}

func main() {
//variables
	x := 5
	var y int = 7
	sum := x + y
//print
	fmt.Println(sum)
//conditonals
	if x > sum {
		fmt.Println("x or y was neg")
	} else {
		fmt.Println("thats pretty good")
	}
//arrays
	a := [5]int{5, 4, 3, 2, 1}
	fmt.Println(a)
//slices
	b := []int{10, 9, 8, 7, 6}
	fmt.Println(b)

//	c := []int{5, 4, 3, 2, 1}
//append 
	b = append(b, 10)

	fmt.Println(b)

//map structure
	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	delete(vertices, "square")

	fmt.Println(vertices)

//indexed for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

//while loop
	i:= 5

	for i > 0{
		fmt.Println(i)
		i--
	}

//while loop with array
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}

//while loop with map
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key", key, "value", value)
	}

//calling a function
	result := sumate(10, 100)
	fmt.Println(result)

	result2, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	}else{
	fmt.Println(result2)
	}
	
	result3, err2 := sqrt(-16)
	if err2 != nil {
		fmt.Println(err2)
	}else{
	fmt.Println(result3)
	}

//using a struct
	p := person{name: "Joshua", age: 21}	
	fmt.Println(p)
	fmt.Println(p.age)
	fmt.Println(p.name)

//pointers
	j := 7
	inc(&j)
	fmt.Println(j)
}

//making a function
func sumate(x int, y int) int {
	return x + y
}

//func with mult returns
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Complex")
	}

	return math.Sqrt(x), nil
}

//func to ++ used in pointers

func inc(x *int) {
	*x++
}