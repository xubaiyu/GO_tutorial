 package main
 
 import "fmt"
 import "unsafe"
 
 const (
    a3 = "abc"
    b3 = len(a3)
    c3 = unsafe.Sizeof(a3)
)

const (
    a4 = iota
    b4 = iota
    c4 = iota
)
 
func main(){
 const b1 string = "abc"
const a1 = "def"
 print(a1,b1)
 fmt.Println(a1,b1)
 
 
  const LENGTH int = 10
   const WIDTH int = 5   
   var area int
   const a, b, c = 1, false, "str" //¶ᗘ¸³ֵ

   area = LENGTH * WIDTH
   fmt.Printf("面积为 : %d", area)
   println()
   println(a, b, c)   
   
   println(a3, b3, c3)   
   println(a4, b4, c4)  
   A()
   B()
   C()
 }
 
 func A(){
  const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
 }
 
 func B(){
 const (
    i=1<<iota
    j=3<<iota
    k
    l
  )


    fmt.Println("i=",i)
    fmt.Println("j=",j)
    fmt.Println("k=",k)
    fmt.Println("l=",l)

 }
 
 func C(){
 
 var a = "helloJOERIGJEGJE;LGKD;KGAL;KSGS;GHK"
 var len = unsafe.Sizeof(a)
 println(len)
 }
