package main

import "fmt"

var a1 = "菜鸟教程"
var b1 string = "runoob.com"
var c1 bool

var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a3 int
    b3 bool
)

var c3, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main() {
   fmt.Println("Hello, World!")
   var a = 1.5
   var b =2
   fmt.Println(a,b)
   fmt.Println(a1,b1,c1)
   println(a1, b1, c1)
   println(x, y, a3, b3, c3, d, e, f)
}