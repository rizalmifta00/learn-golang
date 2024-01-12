package main
import "fmt"

func main(){
	
	var xs = "123" // string
	for i, v := range xs {
		fmt.Println("Index=", i, "Value=", string(v))
	}
	
}