package main

import "fmt"

func main(){

	var point = 2

	if point ==10{
		fmt.Println("Lulus dengan nilai sempurna")
	}else if point > 5 {
		fmt.Println("Lulus")
	}else{
		fmt.Printf("Tidak Lulus, nilai Anda %d\n" ,point )
	}
}