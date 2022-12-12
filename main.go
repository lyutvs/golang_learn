package main

import (
	"fmt"
	"golang_learn/account"
	"golang_learn/mydict"
)

func main() {
	account := account.NewAccount("sehyoun")
	account.Deposit(10)
	fmt.Println(account.Balance())

	//err := account.WithDraw(20)
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println(account.Balance(), account.Owner())

	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("second")

	word := "hello"
	gretingWord := "Greeting"
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	errAdd := dictionary.Add(word, gretingWord)

	if errAdd != nil {
		fmt.Println(errAdd)
	}
	hello, err := dictionary.Search(word)
	fmt.Println(hello)

	erre := dictionary.Update(";s;s;s;s", "Second")

	baseWord := "hellow"

	if err != nil {
		fmt.Println(erre)
	}

	wordd, _ := dictionary.Search(baseWord)

	dictionary.Delete(baseWord)

	fmt.Println(wordd)
}
