package main

import "fmt"

func main(){

	var names [4]string
	names[0]= "james"
	names[1] = "arthur"
	names[2] = "dibala"
	names[3] = "ronaldo"

	

	for _,v:= range names {
		fmt.Print(" ",v)
	}
	fmt.Println() 

}