package main

import "fmt"

func main() {
   var numbers = make([]int,3,5)
   
   var slice1 = []int{1,2,3}
   printSlice(slice1)
   
   slice2:= []int{1,2,3}
   printSlice(slice2)

   printSlice(numbers)
   
   var numbers3 []int

   printSlice(numbers3)

   if(numbers3 == nil){
      fmt.Printf("切片是空的")
   }
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}