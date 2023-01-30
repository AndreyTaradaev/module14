package main

import (
	"fmt"
	"math"
)

func  GetIntSqrt(x uint) uint {
 r := math.Sqrt(float64(x))

return  uint (r) *uint (r)
}


func main(){
var a,b uint
//var b uint

for a=1 ; a <=10000; a++{
	for b = 1; b<=a ;b++{
		s :=a*a+b*b
		if(s == GetIntSqrt(s)) {
			fmt.Println( "a: ",a, "b: ",b, "c:= ",s )
		}
	}
}



}