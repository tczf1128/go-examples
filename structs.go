package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println(Person{"Bob", 24})

	fmt.Println(Person{name: "Alice", age: 23})

	fmt.Println(Person{name: "Fred"})

	fmt.Println(&Person{name: "Ann", age: 20})

	s := Person{name: "Sean", age: 30}
	fmt.Println(s.name)

	s.age = 90
	fmt.Println(s.age)

	sp := &s
	fmt.Println(sp.age)
	fmt.Println((*sp).age)

	sp.age = 51
	fmt.Println(sp.age)

}
