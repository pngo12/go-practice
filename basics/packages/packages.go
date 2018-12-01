package main

// Style of importing multiple packages
import (
	"fmt" 
	"math"
	"go/basics/packages/packages"
)

func main (){
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.3))
	fmt.Println(math.Sqrt(16))
	fmt.printLn(reverse.Reverse("hello"))
}