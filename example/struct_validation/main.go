package main

import (
	"fmt"
	"log"

	"github.com/fn-code/govals"
)

type validation struct {
	Name    string `govals:"len, max=22, min=10"`
	Age     int8   `govals:"len, min=18, max=20"`
	Email   string `govals:"email"`
	Date    string `govals:"date"`
	Times   string `govals:"time"`
	Hoby    string `govals:"alpha"`
	Max     string `govals:"numeric"`
	Text    string `govals:"alphaNumeric"`
	Phone   string `govals:"phone"`
	Address string `govals:"custom, regex=[^(a-zA-Z)]"`
}

func main() {
	valid := &validation{
		Name:    "Ludin Ne",
		Age:     18,
		Email:   "lud!_${}}#yyngmail.com",
		Address: "Jl.Irigasi",
		Times:   "12:00:01",
		Date:    "12-31-1996",
		Hoby:    "Music",
		Max:     "233330",
		Phone:   "+6282290102726",
		Text:    "Ludin123--s",
	}

	gov := govals.Struct(valid)
	ok, err := gov.GetStatus("Text")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ok)

}
