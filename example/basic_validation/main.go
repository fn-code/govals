package main

import (
	"fmt"
	"log"

	"github.com/fn-code/govals"
)

func main() {
	ok, err := govals.Length(1000, 2, 10)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ok)
	ok, err = govals.RuneLength("ludin nento", 20, 30)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ok)

	ok, err = govals.Email("ludynn77654.com")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ok)

}
