package main

import (
	"fmt"
	"golang_learn/somting"
	"strings"
)

type person struct {
	name    string
	age     int
	favFood []string
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(word ...string) {
	fmt.Println(word)
}

func lenAndUppers(name string) (length int, upper string) {
	defer fmt.Println("bybyebyebeye")
	length = len(name)
	upper = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	defer fmt.Println("superAdd")
	for index, num := range numbers {
		fmt.Println(index, num)
	}
	return 1
}

func loopAdd(numbers ...int) int {
	defer fmt.Println("loopAdd")
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func naturalLoop(number ...int) int {
	for i := 0; i < len(number); i++ {
		fmt.Println(number[i])
	}
	return 1
}

func canIDrive(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func canIdDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrive(18))

	fmt.Println(canIdDrink(16))

	totalLength, upperName := lenAndUpper("hello")

	fmt.Println(totalLength, upperName)

	totalLengths, upperNames := lenAndUppers("jeonsehyoun")
	fmt.Println(totalLengths, upperNames)

	repeatMe("hello", "hohhooh", "qwdqw", "12343")
	somting.Bye()

	superAdd(1, 2, 3, 4, 5, 6)

	naturalLoop(1, 2, 3, 4, 5, 6, 7, 8, 9)

	result := loopAdd(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(result)

	names := []string{"nico", "lynn", "dal"}
	names = append(names, "flynn")
	fmt.Println(names)

	sehyoun := map[string]string{"name": "sehyoun", "age": "12"}
	fmt.Println(sehyoun)

	favFood := []string{"kimchi", "ramen"}
	sehyounn := person{"sehyoun", 18, favFood}
	sehyonType := person{name: "sehyoun", age: 18, favFood: favFood}
	fmt.Println(sehyounn)
	fmt.Println(sehyonType)
}
