package main

import "fmt"

func main(){

	// tidak boleh ada variabel yang tidak di pakai

	var first,second,third string 
	first, second, third = "satu","dua","tiga"

	fmt.Printf("%s %s %s \n", first,second,third)

	var four, five , six string = "empat","lima","enam";

	fmt.Printf("%s %s %s \n",four,five,six)

	seven, eight, nine := "tujuh","delapan","sembilan"

	fmt.Printf("%s %s %s  \n" , seven,eight,nine)
}