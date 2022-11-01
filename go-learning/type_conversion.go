///convert interger to float
package main

import "fmt"

func main(){
	var number int =  34
	fmt.Printf("Data =%v,Type = %T\n" ,number,number);

	var fnumber float32 = float32(number)
	fmt.Printf("Data = %v,Type = %T",fnumber,fnumber);
}