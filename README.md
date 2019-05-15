# Govals
Govals is a go package for validate data, using struct and etc. govals have custom function like validate email, numeric, alpha, alpha numeric, date, time, phone number, and length of int or string, i hope soon i will add more function for validate data.
# Example
example how to use govals, you can open folder [example](https://github.com/fn-code/govals/tree/master/example)
# Usage
Run this command `go get github.com/fn-code/govals` to use this library
- ### Basic Functions
    this is basic function you can use to validate your data:
    ```go
    ok, err := govals.Length(1000, 2, 10)
	if err != nil {
        log.Println(err)
    }
    if !ok {
        // bla bla bla
    }
    ```
- ### Struct
    in struct you can use `govals` tag like this:
    ```go
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
    ```
    - `len, min=10, max=22` validate length of string or int
    - `email` validate string of email
    - `date` validate date of string if match, like 09-12-1996 or 09/12/1996
    - `time` validate time if match, like 12:00:01 or 12:01
    - `alpha`validate string if contains alpha character
    - `numeric`validate string if contains numeric
    - `alphaNumeric` validate string if contains alpha numeric character
    - `phone` validate string if valid phone number format
    - `custom` validate custom rule, using regexp, you can write your own rule to validate data

