package main

import "fmt"

func main(){
	dataSlice := []int{0,1,2,3,4,5,6,7,8,9,10}
	//fmt.Println(dataSlice)
	for _, i := range dataSlice{
		if i % 2 == 0 {
			fmt.Println(i," is even")
		} else{
			fmt.Println(i," is odd")
		}
	}
}