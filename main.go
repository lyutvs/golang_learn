package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [2]string{"sehyoun", "qdqdw"}
	for _, person := range people {
		go isSexy(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

	go sexyCount("sehyoun")
	go sexyCount("adwa")
	time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "Is Sexy", i)
		time.Sleep(time.Second)
	}
}
func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- person + "is sexy"
}
