package main
import "fmt"

func main(){
	
	// var xs = "123" // string
	// for i, v := range xs {
	// 	fmt.Println("Index=", i, "Value=", string(v))
	// }
	var total = 0
	var ys = [5]int{10,20,30,40,50}
	for _,v:=range ys {

		total +=  v
	}
	
	fmt.Println(total)
}