学习笔记

remind
In Go, panic is not exception


《Go进阶训练营第四课问题收集》 https://shimo.im/docs/R6gP8qyvWqJrgRCk/



```go
# recover 野生 goroutine
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


* Sentinel error: 耦合太大
* Opaque error: 老师认为是最灵活的错误处理策略，要求代码和调用者之间的耦合最少
* Wrap errors：打印堆栈信息
* 最佳实践： pkg/errors
* go1.13: 不够完整
* go2：可以期待下
    
