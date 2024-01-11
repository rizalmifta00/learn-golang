package main

import "fmt"

func main(){
	var point = 3

	switch point {
	case 8 :
		fmt.Println("perfect")
	case 7,6,5,4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can better!")
		}
	}
}