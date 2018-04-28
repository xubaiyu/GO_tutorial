package main
import "fmt"


type Integer int
func (a Integer) Less(b Integer) bool {
    return a < b
}

type Integer1 int

func Integer_Less(a Integer1, b Integer1) bool {
return a < b
}


func main() {
    var a Integer = 1
    if a.Less(2) {
        fmt.Println(a, "Less 2")
    }
    var a1 Integer1 = 1
    if Integer_Less(a1, 2) {
    fmt.Println(a1, "Less 2")
    
    var a5 int = 10   

    fmt.Printf("变量的地址: %x\n", &a5  )
    }
    
    A()
    B()
}

func A(){
  var a int= 20   /* 声明实际变量 */
   var ip *int        /* 声明指针变量 */

   ip = &a  /* 指针变量的存储地址 */

   fmt.Printf("a 变量的地址是: %x\n", &a  )

   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip )
   var  ptr *int

   fmt.Printf("ptr 的值为 : %x\n", ptr  )
}

func B(){
type Books struct {
   title string
   author string
   subject string
   book_id int
}

var Book1 Books        /* 声明 Book1 为 Books 类型 */
   var Book2 Books        /* 声明 Book2 为 Books 类型 */

   /* book 1 描述 */
   Book1.title = "Go 语言"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go 语言教程"
   Book1.book_id = 6495407

   /* book 2 描述 */
   Book2.title = "Python 教程"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python 语言教程"
   Book2.book_id = 6495700

   /* 打印 Book1 信息 */
   fmt.Printf( "Book 1 title : %s\n", Book1.title)
   fmt.Printf( "Book 1 author : %s\n", Book1.author)
   fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
   fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

   /* 打印 Book2 信息 */
   fmt.Printf( "Book 2 title : %s\n", Book2.title)
   fmt.Printf( "Book 2 author : %s\n", Book2.author)
   fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
   fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)

}