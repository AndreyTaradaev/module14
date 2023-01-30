package main

import (
	"fmt"
)


type checked struct{
	ar1, ar2 bool
}

func (c checked) Haveboth() bool {
return  c.ar1&&c.ar2
} 

 func  (c checked) IsBoth()  bool {		
	return  c.ar1 && c.ar2
}  

func (c checked) String() string{
return fmt.Sprintf( "{ %v, %v }", c.ar1,c.ar2 )	
}

func CreateChecked( c *checked, x1,x2 bool) *checked {
	if(c==nil){
		c = new(checked)	
	}
	if(x1==true){
		c.ar1 = x1
	}
	if(x2==true){
		c.ar2 = x2
	}
return c
}

func GetArray( x map[int] *checked )  []int  {	
	ar := make([]int,0,len(x))
	for key, value := range x {
		if( value.IsBoth()) {
		ar = append(ar, key)	
		}
	}
return ar
}



func main (){
m:= make(map[int] *checked)	
var ui1,ui2 uint
var in int
fmt.Println("Enter first array size:")	
fmt.Scanln(&ui1)
if(ui1==0){
fmt.Println("first array  is empty")
return
}

fmt.Println("Enter second array size:")	
fmt.Scanln(&ui2)
if(ui2==0){
fmt.Println("second array  is empty")
return
}

fmt.Println(" Fill in first array")

for i:=uint(0); i<ui1 ;i++{
	fmt.Scanln(&in)	 
	m[in] = CreateChecked(nil,true,false)
}
fmt.Println(" Fill in second array")

for i:=uint(0); i<ui2 ;i++{
	fmt.Scanln(&in)	
	if value ,inMap := m[in] ; inMap {
		value.ar2 = true
		continue
	}
	m[in] = CreateChecked(nil,false,true)
}
fmt.Println(m)
fmt.Println( GetArray (m))
fmt.Println( GetArray (m))
}