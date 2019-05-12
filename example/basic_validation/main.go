package main

import (
	"fmt"
	"log"

	"github.com/fn-code/govals"
)

func main() {
	val, err := govals.Length(1000, 2, 10)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(val)
	val2, err := govals.RuneLength("ludin nento", 20, 30)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(val2)

	val3, err := govals.Email("ludynn77654.com")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(val3)

}
