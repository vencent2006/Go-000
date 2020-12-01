学习笔记

remind
In Go, panic is not exception


《Go进阶训练营第四课问题收集》 https://shimo.im/docs/R6gP8qyvWqJrgRCk/

```go
package main
import(
  "fmt"
  "time"
)

func main(){
  fmt.Println("vim-go")
  Go(func(){
    fmt.Println("hello")
    //panic
    panic("一路向西")
  })
  
  time.Sleep(5*time.Second)

}

func Go(x func()){
  go func(){
    defer func(){
      if err := recover(); err != nil{
        fmt.Println(err)
      }
    }()
    
    //exec x()
    x()
    
  }()
}

```
