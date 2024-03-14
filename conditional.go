package main

import "fmt"

func main() {
	//if statement
	var userName = "Joko"
	if userName == "Joko" {
		fmt.Println("benar")
	}
	// if else statement
	var hindu string = "david"
	if hindu == "akmal" {
		fmt.Println("itu memang benar")
	} else {
		fmt.Println("baru saja murtad")
	}

	//else if else if statement
	var umur int = 12
	if umur >= 12 {
		fmt.Println("benar")
	} else if umur <= 10 && umur >= 10 {
		fmt.Println("salah")
	} else {
		fmt.Println("tidak tahu")
	}
}
